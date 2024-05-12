package main

import (
	"fmt"
	"rpgGo/rpgGo"
)

func main() {

	fmt.Println("Welcome to GoLang Island!")

	fmt.Println("What is your name?")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)

	classPick := rpgGo.PickClass()
	fmt.Println(rpgGo.BuildCharacter(name, classPick))
}
