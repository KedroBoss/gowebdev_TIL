package main

import (
	"fmt"
)

// Simple example of using custom variables in html
func main() {
	name := "Boss"

	tmp := `
	<DOCTYPE html>
	<html>
	<head>
	</head>
	<body>
	<h1>Hello ` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tmp)
}
