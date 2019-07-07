package main

import (
	"fmt"
	"id-generator/generator"
)

func main() {

	// Example: this will give us a 44 byte, base64 encoded output
	token, _ := generator.GenerateID(24, "usr")
	fmt.Print(token)
}
