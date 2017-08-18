# Templates
Generic letter, will be customized and personalized

Hello <username>

stdlib has 2 packages:  

* text/template - foundation for the second
* html/templates

go run main.go > index.html

### text/template

.gohtml - custom file extension, not mandatory, can be changed to any extension. I'll be using html, just because code editors can use appropriate tools

```tmp := template.ParseFiles(...files)``` - creates a container holding templates

```tmp = tmp.ParseFiles(...files)``` - adds templates to the container

tmp.ParseGlob

Must() does error checking