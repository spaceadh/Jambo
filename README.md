<h1 align="center">Jambo🔥PROGRAMMING🔥LANGUAGE</h1>
<p align="center">
    <a href="https://github.com/spaceadh/Jambo"><img alt="Jambo Programming Language" src="https://img.shields.io/badge/Jambo-Programming%20Language-yellow"></a>
    <a href="https://github.com/spaceadh/Jambo"><img alt="Jambo Programming Language" src="https://img.shields.io/badge/platform-Linux | Windows | Android-green.svg"></a>
    <a href="https://github.com/spaceadh/Jambo"><img alt="Jambo Programming Language" src="https://img.shields.io/github/last-commit/spaceadh/Jambo"></a>
<br>
    <a href="https://github.com/spaceadh/Jambo"><img alt="Jambo Programming Language" src="https://img.shields.io/github/downloads/avicennajr/jambo/total"></a>
    <a href="https://github.com/spaceadh/Jambo/releases"><img alt="Jambo Programming Language" src="https://img.shields.io/github/v/release/avicennajr/jambo?include_prereleases"></a>
    <a href="https://github.com/spaceadh/Jambo"><img alt="Jambo Programming Language" src="https://img.shields.io/github/actions/workflow/status/spaceadh/Jambo/tests.yml?style=plastic"></a>
<br>
    <a href="https://github.com/spaceadh/Jambo"><img alt="Jambo Programming Language" src="https://img.shields.io/github/stars/spaceadh/Jambo?style=social"></a>
</p>
A Swahili Programming Language of its kind built from the ground up.

## Installation

To get started download the executables from the release page or follow the instructions for your device below:

### Linux

 - Download the binary:

```
curl -O -L https://github.com/spaceadh/Jambo/releases/download/v0.5.16/jambo_Linux_amd64.tar.gz
```

  - Extract the file to make global available:

```
sudo tar -C /usr/local/bin -xzvf jambo_Linux_amd64.tar.gz
```

 - Confirm installation with:

```
jambo -v
```


### MacOs ( Apple silicon Mac )

 - Download the binary:

    - For apple silicon mac use:

        ```
        curl -O -L https://github.com/spaceadh/Jambo/releases/download/v0.5.16/jambo_Darwin_arm64.tar.gz
        ```

    - For apple intel mac use:

        ```
        curl -O -L https://github.com/spaceadh/Jambo/releases/download/v0.5.16/jambo_Darwin_amd64.tar.gz
        ```

- Extract the file to make global available:
    - For apple silicon mac use:

        ```
        sudo tar -C /usr/local/bin -xzvf jambo_Darwin_arm64.tar.gz
        ```
    - For apple intel mac use:

        ```
        sudo tar -C /usr/local/bin -xzvf jambo_Darwin_amd64.tar.gz
        ```

- Confirm installation with:

```
jambo -v
```


### Android (Termux)

 - Make sure you have [Termux](https://f-droid.org/repo/com.termux_118.apk) installed.
 - Download the binary with this command:

```
curl -O -L https://github.com/spaceadh/Jambo/releases/download/v0.5.16/jambo_Android_arm64.tar.gz
```
 - Extract the file:

```
tar -xzvf jambo_Android_arm64.tar.gz
```
 - Add it to path:

```
echo "alias jambo='~/jambo'" >> .bashrc
```
 - Confirm installation with:

```
jambo -v 
```

### Windows

- Executable:
    - Download the Jambo zip file [Here](https://github.com/spaceadh/Jambo/releases/download/v0.5.16/jambo_Windows_amd64.zip)
    - Unzip to get the executable
    - Double click the executable

- Jambo Installer
    > Coming Soon

### Building From Source

 - Make sure you have golang installed (atleast 1.19.0 and above)
 - Run the following command:

```
go build -o jambo main.go
```
 - You can optionally add the binary to $PATH as shown above
 - Confirm installtion with:

```
jambo -v
```

## Syntax At A Glance

**NOTE**
> There is a more detailed documentation of the language [here](https://jamboprogramming.org).

Jambo, although still in its early stage, intends to be a fully functional programming language, and thus it has been baked with many features.

### Defining A Variable

You can define variables like this:

```
x = 2;
y = 3;

andika(x*y) // output is 6
```
You can also use the `fanya` keyword to define a variabe:
```
fanya x = 3
```
**Note that `fanya` keyword is OPTIONAL**

### Comments

Jambo supports both single line and multiple line comments as shown below:

```
// Single line comment

/*
Multiple
Line
Comment 
*/ 
```

### Arithmetic Operations

For now Jambo supports `+`, `-`, `/`, `*` and `%`. Jambo also provides precedence of operations using the BODMAS rule:

```
2 + 2 * 3 // output = 8

2 * (2 + 3) // output = 10
```

### Types

Jambo has the following types:

Type      | Syntax                                    | Comments
--------- | ----------------------------------------- | -----------------------
BOOL      | `kweli sikweli`                           | kweli == true, sikweli == false
INT       | `1, 100, 342, -4`                         | These are signed 64 bit integers
FLOAT     | `2.3, 4.5. 100.8094`                      | Signed 64 bit floats
STRING    | `"" "mambo" "habari yako"`                | They can be in double `"` or single `'` quotes
ARRAY     | `[] [1, 2, 3] [1, "moja", kweli]`         | Arrays can hold any types
DICT      | `{} {"a": 3, 1: "moja", kweli: 2}`        | Keys can be int, string or bool. Values can be anything
NULL      | `tupu`                                    | These are nil objects

### Functions

This is how you define a function in Jambo:

```
jumlisha = unda(x, y) {
        rudisha x + y
    }

andika(jumlisha(3,4))
```

Jambo also supports recursion:

```
fibo = unda(x) {
	kama (x == 0) {
		rudisha 0;
	} au kama (x == 1) {
			rudisha 1;
	} sivyo {
			rudisha fibo(x - 1) + fibo(x - 2);
	}
}
```

### If Statements

Jambo supports if, elif and else statements with keywords `kama`, `au kama` and `sivyo` respectively:

```
kama (2<1) {
    andika("Mbili ni ndogo kuliko moja")
} au kama (3 < 1) {
    andika ("Tatu ni ndogo kuliko moja")
} sivyo {
    andika("Moja ni ndogo")
}
```

### While Loops

Jambo's while loop syntax is as follows:

```
i = 10

wakati (i > 0) {
	andika(i)
	i--
}
```

### Arrays

This is how you initiliaze and perform other array operations in Jambo:
```
arr = []

// To add elements

sukuma(arr, 2)
andika(arr) // output = [2]
// Add two Arrays

arr2 = [1,2,3,4]

arr3 = arr1 + arr2

andika(arr3) // output = [2,1,2,3,4]

// reassign value

arr3[0] = 0

andika[arr3] // output = [0,1,2,3,4]

// get specific item

andika(arr[3]) // output = 3
```

### Dictionaries

Jambo also supports dictionaries and you can do a lot with them as follows:
```
mtu = {"jina": "Mojo", "kabila": "Mnyakusa"}

// get value from key 
andika(mtu["jina"]) // output = Mojo

andika(mtu["kabila"]); // output = Mnyakusa

// You can reassign values

mtu["jina"] = "Victor Alvin"

andika(mtu["jina"]) // output = Victor Alvin

// You can also add new values like this:

mtu["anapoishi"] = "Dar Es Salaam"

andika(mtu) // output = {"jina": "Victor Alvin", "kabila": "Mnyakusa", "anapoishi": "Dar Es Salaam"}

// You can also add two Dictionaries

kazi = {"kazi": "jambazi"}

mtu = mtu + kazi

andika(mtu) // output = {"jina": "Victor Alvin", "kabila": "Mnyakusa", "anapoishi": "Dar Es Salaam", "kazi": "jambazi"}
```

### For Loops

These can iterate over strings, arrays and dictionaries:
```
kwa i ktk "habari" {
    andika(i)
}
/* //output
h
a
b
a
r
i
*/
```

### Getting Input From User

In Jambo you can get input from users using the `jaza()` keyword as follows:
```
jina = jaza("Unaitwa nani? ") // will prompt for input

andika("Habari yako " + jina)
```

## How To Run

### Using The Intepreter:

You can enter the intepreter by simply running the `jambo` command:
```
jambo
>>> andika("karibu")
karibu
>>> 2 + 2
4
```
Kindly Note that everything should be placed in a single line. Here's an example:
```
>>> kama (x > y) {andika("X ni kubwa")} sivyo {andika("Y ni kubwa")}
```
### Running From File

To run a Jambo script, write the `jambo` command followed by the name of the file with a `.jb` or `.sw` extension:

```
jambo myFile.jb
```

## Issues

Kindly open an [Issue](https://github.com/spaceadh/Jambo/issues) to make suggestions and anything else.

## Contributions

### Documentation

There are documentations for two languages, English and Kiswahili, which are both under the `docs` folder. All files are written in markdown. Feel free to contribute by making a pull request.

### Code

Clone the repo, hack it, make sure all tests are passing then submit a pull request.

 > Make sure ALL tests are passing before making a pull request. You can confirm with running `make tests`

## Community

Jambo has a passionate community, join us on [Telegram](https://t.me/JamboProgrammingChat)

## License

[MIT](http://opensource.org/licenses/MIT)

## Authors

Jambo Programming Language has been authored and being actively maintained by [Victor Alvin](https://github.com/spaceadh)