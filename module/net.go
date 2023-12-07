package module

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt" 
	"io"
	"os"
	"strings" 
	"mime"
	"os/user"
    "path/filepath"

	"github.com/spaceadh/Jambo/object"
)

var NetFunctions = map[string]object.ModuleFunction{}

func init() {
	NetFunctions["peruzi"] = getRequest
	NetFunctions["tuma"] = postRequest
	NetFunctions["buffer"] = generateAuth
	NetFunctions["password"] = password
	NetFunctions["wget"] = wget
}
// generateAuth is a function that generates an authorization string using consumerKey and consumerSecret.
func generateAuth(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "generateAuth function requires exactly two arguments: consumerKey and consumerSecret"}
	}

	consumerKey, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "generateAuth function requires a string argument for consumerKey"}
	}

	consumerSecret, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "generateAuth function requires a string argument for consumerSecret"}
	}

	// Concatenate consumerKey and consumerSecret with a colon
	authString := consumerKey.Value + ":" + consumerSecret.Value

	// Encode the concatenated string to base64
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(authString))

	return &object.String{Value: encodedAuth}
}

// password is a function that generates an authorization string using businessCode, passkey, and timestamp.
func password(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 3 {
		return &object.Error{Message: "password function requires exactly three arguments: businessCode, passkey, and timestamp"}
	}

	businessCode, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "password function requires a string argument for businessCode"}
	}

	passkey, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "password function requires a string argument for passkey"}
	}

	timestamp, ok := args[2].(*object.String)
	if !ok {
		return &object.Error{Message: "password function requires a string argument for timestamp"}
	}

	// Concatenate businessCode, passkey, and timestamp with a colon using the + operator
	authString := businessCode.Value + "+" + passkey.Value + "+" + timestamp.Value

	// Encode the concatenated string to base64
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(authString))

	return &object.String{Value: encodedAuth}
}

func getRequest(args []object.Object, defs map[string]object.Object) object.Object {

	if len(defs) != 0 {
		var url *object.String
		var headers, params *object.Dict
		for k, v := range defs {
			switch k {
			case "yuareli":
				strUrl, ok := v.(*object.String)
				if !ok {
					return &object.Error{Message: "Yuareli iwe neno"}
				}
				url = strUrl
			case "vichwa":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Vichwa lazima viwe kamusi"}
				}
				headers = dictHead
			case "mwili":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Mwili lazima iwe kamusi"}
				}
				params = dictHead
			default:
				return &object.Error{Message: "Hoja si sahihi. Tumia yuareli na vichwa."}
			}
		}
		if url.Value == "" {
			return &object.Error{Message: "Yuareli ni lazima"}
		}

		var responseBody *bytes.Buffer
		if params != nil {
			booty := convertObjectToWhatever(params)

			jsonBody, err := json.Marshal(booty)

			if err != nil {
				return &object.Error{Message: "Huku format query yako vizuri."}
			}

			responseBody = bytes.NewBuffer(jsonBody)
		}

		var req *http.Request
		var err error
		if responseBody != nil {
			req, err = http.NewRequest("GET", url.Value, responseBody)
		} else {
			req, err = http.NewRequest("GET", url.Value, nil)
		}
		if err != nil {
			return &object.Error{Message: "Tumeshindwa kufanya request"}
		}

		if headers != nil {
			for _, val := range headers.Pairs {
				req.Header.Set(val.Key.Inspect(), val.Value.Inspect())
			}
		}
		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {
			return &object.Error{Message: "Tumeshindwa kutuma request."}
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &object.Error{Message: "Tumeshindwa kusoma majibu."}
		}

		return &object.String{Value: string(respBody)}

	}

	if len(args) == 1 {
		url, ok := args[0].(*object.String)
		if !ok {
			return &object.Error{Message: "Yuareli lazima iwe neno"}
		}
		req, err := http.NewRequest("GET", url.Value, nil)
		if err != nil {
			return &object.Error{Message: "Tumeshindwa kufanya request"}
		}

		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {
			return &object.Error{Message: "Tumeshindwa kutuma request."}
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &object.Error{Message: "Tumeshindwa kusoma majibu."}
		}

		return &object.String{Value: string(respBody)}
	}
	return &object.Error{Message: "Hoja si sahihi. Tumia yuareli na vichwa."}
}

func postRequest(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		var url *object.String
		var headers, params *object.Dict
		for k, v := range defs {
			switch k {
			case "yuareli":
				strUrl, ok := v.(*object.String)
				if !ok {
					return &object.Error{Message: "Yuareli iwe neno"}
				}
				url = strUrl
			case "vichwa":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Vichwa lazima viwe kamusi"}
				}
				headers = dictHead
			case "mwili":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Mwili lazima iwe kamusi"}
				}
				params = dictHead
			default:
				return &object.Error{Message: "Hoja si sahihi. Tumia yuareli na vichwa."}
			}
		}
		if url.Value == "" {
			return &object.Error{Message: "Yuareli ni lazima"}
		}
		var responseBody *bytes.Buffer
		if params != nil {
			booty := convertObjectToWhatever(params)

			jsonBody, err := json.Marshal(booty)

			if err != nil {
				return &object.Error{Message: "Huku format query yako vizuri."}
			}

			responseBody = bytes.NewBuffer(jsonBody)
		}
		var req *http.Request
		var err error
		if responseBody != nil {
			req, err = http.NewRequest("POST", url.Value, responseBody)
		} else {
			req, err = http.NewRequest("POST", url.Value, nil)
		}
		if err != nil {
			return &object.Error{Message: "Tumeshindwa kufanya request"}
		}
		if headers != nil {
			for _, val := range headers.Pairs {
				req.Header.Set(val.Key.Inspect(), val.Value.Inspect())
			}
		}
		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {
			return &object.Error{Message: "Tumeshindwa kutuma request."}
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &object.Error{Message: "Tumeshindwa kusoma majibu."}
		}
		return &object.String{Value: string(respBody)}
	}
	return &object.Error{Message: "Hoja si sahihi. Tumia yuareli na vichwa."}
}

// wget downloads a file from the specified URL and returns the path where it is saved.
func wget(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "wget function requires exactly one argument: url"}
	}

	url, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "wget function requires a string argument for the URL"}
	}

	// Perform the GET request
	req, err := http.NewRequest("GET", url.Value, nil)
	if err != nil {
		return &object.Error{Message: "Failed to create a request"}
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return &object.Error{Message: "Failed to send GET request"}
	}
	defer resp.Body.Close()

	// Extract filename from URL or Content-Disposition header
	fileName := getFileName(url.Value, resp.Header)

	// Get the current user
	currentUser, err := user.Current()
	if err != nil {
		return &object.Error{Message: "Failed to determine the user's home directory"}
	}

	// Construct the file path in the Downloads directory
	downloadPath := filepath.Join(currentUser.HomeDir, "Downloads", fileName)

	fmt.Printf("Downloading: %s\n", url.Value)
	fmt.Printf("Content Length: %s\n", resp.Header.Get("Content-Length"))
	fmt.Printf("Content Type: %s\n", determineFileType(resp.Header, resp.Body))

	// Create a file to save the downloaded content in the Downloads directory
	file, err := os.Create(downloadPath)
	if err != nil {
		return &object.Error{Message: "Failed to create a file for download"}
	}
	defer file.Close()

	// Copy the response body to the file with progress
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return &object.Error{Message: "Failed to download content"}
	}

	fmt.Printf("\nDownload complete. File saved as: %s\n", downloadPath)

	return &object.String{Value: downloadPath}
}

// getFileName extracts the filename from the URL or Content-Disposition header.
func getFileName(url string, header http.Header) string {
	// Try to extract filename from the URL
	_, file := filepath.Split(url)

	// If filename not found in URL, try to extract from Content-Disposition header
	if file == "" {
		file = getFileNameFromHeader(header)
	}

	// If both fail, use a default name
	if file == "" {
		file = "JamboDownloads"
	}

	return file
}

// getFileNameFromHeader extracts the filename from the Content-Disposition header.
func getFileNameFromHeader(header http.Header) string {
	contentDisposition := header.Get("Content-Disposition")
	if contentDisposition != "" {
		_, params, err := mime.ParseMediaType(contentDisposition)
		if err == nil {
			filename, ok := params["filename"]
			if ok {
				return filename
			}
		}
	}

	// If the filename is not found in Content-Disposition, use a default name
	return ""
}

// determineFileType attempts to determine the file type based on the content itself.
func determineFileType(header http.Header, body io.Reader) string {
	// Try to get content type from header
	contentType := header.Get("Content-Type")

	// If content type is not provided or is "application/octet-stream" (default binary type),
	// attempt to determine the file type based on the first few bytes of the content.
	if contentType == "application/octet-stream" || contentType == "" {
		buf := make([]byte, 512) // Read the first 512 bytes

		n, err := body.Read(buf)
		if err != nil && err != io.EOF {
			return "unknown"
		}

		// Reset the reader to allow further reading
		body = io.MultiReader(strings.NewReader(string(buf[:n])), body)

		contentType = http.DetectContentType(buf)
	}

	return contentType
}