package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Simple example of using custom variables in html
func main() {
	name := "Boss"

	tmp := fmt.Sprint(`
	<DOCTYPE html>
	<html>
	<head>
	</head>
	<body>
	<h1>Hello ` + name + `</h1>
	</body>
	</html>
	`)
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(tmp))
}
