<h1 align="center">Jambo🔥PROGRAMMING🔥LANGUAGE</h1>

# Jambo Programming Language Extension

![Jambo Extension](https://images2.imgbox.com/f9/cf/OZJTfmlM_o.jpg)

A Kenyan Swahili Programming Language of its kind built from the ground up.

## Overview

The Jambo Programming Language extension for Visual Studio Code provides enhanced support for Jambo files, including syntax highlighting, code completion, and more.

## Syntax At A Glance

**NOTE**

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

### Kuunganisha na Sarufi API
Mahitaji(Requirements):

Kabla ya kutumia Sarufi API na Jambo, hakikisha unavyo:
```
    Vyeti vya Sarufi API (client_id na client_secret)
    Lugha ya programu ya Jambo imewekwa
```
Usanidi

Unda faili ya usanidi (kwa mfano, config.jb) na vyeti vyako vya Sarufi API:
```
tumia mtandao
tumia jsoni
pakeji sarufi {
    andaa = unda(file) {
        config = fungua(file)
        configString = config.soma()
        configDict = jsoni.dikodi(configString)
        clientID = configDict["client_id"]
        clientSecret = configDict["client_secret"]
        params = {"client_id": clientID, "client_secret": clientSecret}
        tokenString = mtandao.tuma(yuareli="https://api.sarufi.io/api/access_token", mwili=params)
        tokenDict = jsoni.dikodi(tokenString)
        @.token = tokenDict["access_token"]
        @.Auth = "Bearer " + @.token
        }

    tokenYangu = unda() {
            rudisha @.token
        }

    tengenezaChatbot = unda(data) {
            majibu = mtandao.tuma(yuareli="https://api.sarufi.io/chatbot", vichwa={"Authorization": @.Auth}, mwili = data)
            rudisha majibu
        }

    pataChatbotZote = unda() {
            majibu = mtandao.peruzi(yuareli="https://api.sarufi.io/chatbots", vichwa={"Authorization": @.Auth})
            rudisha majibu
        }
}

```

## License

[MIT](http://opensource.org/licenses/MIT)

## Authors

Jambo Programming Language has been authored and being actively maintained by [Victor Alvin Wachira ](https://github.com/spaceadh)

### Installation

1. Open Visual Studio Code.
2. Go to Extensions.
3. Search for "Jambo" and install the extension.

Happy coding with Jambo!



