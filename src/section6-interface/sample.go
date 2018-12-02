package main

import (
	"fmt"
)

type Human interface {
	Say() string
	Lgtm()
}

type Person struct {
	Name string
}

func (p Person) Say() string {
	name := "Mr." + p.Name
	fmt.Println(name)
	return name
}

func (p Person) Lgtm() {
	fmt.Println("LGTM!!")
}

func isMike(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
}

func main() {
	var mike Human = Person{"Mike"}
	mike.Say()
	mike.Lgtm()
	isMike(mike)
}
