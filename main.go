package main

import (
	"fmt"
	"rpgGo/rpgGo"
)

func FinishCharacter(name map[string]stri, Stats map[string]rpgGo.Character, Life map[string]rpgGo.Character, Items map[string]rpgGo.Items, Level map[string]rpgGo.PlayerExp) {

}

func main() {

	fmt.Println("Welcome to GoLang Island!")

	fmt.Println("What is your name?")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)

	classPick := rpgGo.PickClass()
	Character, Stats, HpM, Items, Level := rpgGo.BuildCharacter(name, classPick)
	// levelInfo := rpgGo.BuildLevelInfo()
	FinishCharacter(Character, Stats, HpM, Items, Level)

}
