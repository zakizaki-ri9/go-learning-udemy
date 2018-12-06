package mylib

import "fmt"

var Public string = "Public"
var private string = "Private"

type Person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("Human!")
}
