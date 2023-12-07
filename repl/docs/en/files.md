# Files in Nuru

Nuru's ability to deal with files is primitive, and as for now it only allows you to read contents of a file.

## Opening a File

You open a file with the `fungua` keyword. This will return an object of type `FAILI`:
```
fileYangu = fungua("file.txt")

aina(fileYangu) // FAILI
```

## Reading a File

Once you have a file object you can read its contents with the `soma()` method. This will return the contents of the file as a string:

```
fileYangu = fungua("file.txt")

fileYangu.soma()
```

### Writing a File

Create a new file object with the `andikaupya` method. This will require two parameters filename and contents of the file as a string:

```
NewFileName = "new.txt"
newContent = "This is the new content."

andikaupya(NewFileName, newContent)

```