# Templates
Generic letter, will be customized and personalized

Hello $username

stdlib has 2 packages:  

* text/template - foundation for the second
* html/templates

go run main.go > index.html

## text/template

.gohtml - custom file extension, not mandatory, can be changed to any extension. I'll be using html, just because code editors can use appropriate tools

### Parsing files 

```tmp := template.ParseFiles(...files)``` - creates a container holding templates

```tmp = tmp.ParseFiles(...files)``` - adds templates to the container

tmp.ParseGlob

Must() does error checking

Parse -> Execute

### Sending data

```{{.}}``` - the current piece of data

When passing struct, make sure the data is exported(starts with upper case)

Passing function to templates is good to modify data. Functions must be passed before the template is parsed.

Time formating depends on the number that are typed in.

Pipelining - passing value from one function to another

To pass template:
```
{{define "templateName"}}
{{end}}

{{template "templateName"}}
```

Methods(functions with receivers) can be used in a template

## State

State - persistent awareness of who is communicating with whom 

## Enctype

When a form allows to upload a file: use enctype = "multipart/form-data"

## Sessions

Server assigns uuid to a client, which then used to have a state and persist data

Assign uuid to cookie -> take the id from the cookie

### UUID



### Tools

* HMAC - encoder. Provided a secret key, generate a hash. If the input is the same-out is the same.
* BASE64 Encoding - if needs to store in cookie something unusuall(like double quotes), can encode it and store it that way, then decode.
* Web Storage: local storage; session storage; - can use cookies, local storage and session storage to store data on a clients machine.
* Context - can store some data in a request. Usually better to pass some "per request variables" like session id.


## MongoDB - NoSQL
###Database --> Collections --> Document

MongoDB stores Databases. Databases store Collections. Collections store Documents.

