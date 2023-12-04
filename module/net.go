package module

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/spaceadh/Jambo/object"
)

var NetFunctions = map[string]object.ModuleFunction{}

func init() {
	NetFunctions["peruzi"] = getRequest
	NetFunctions["tuma"] = postRequest
	NetFunctions["buffer"] = generateAuth
	NetFunctions["password"] = password
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
