package main

import (
	"fmt"
	"rpgGo/rpgGo"
)

type CharacterInfo struct {
	Name  []string
	Stats map[string]int
	Life  map[string]int
	Items map[string]int
	Level map[string]int
}

func FinishCharacter(name []string, Stats map[string]int, Life map[string]int, Items map[string]int, Level map[string]int) CharacterInfo {

	return CharacterInfo{
		Name:  name,
		Stats: Stats,
		Life:  Life,
		Items: Items,
		Level: Level,
	}
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
	theGuy := FinishCharacter(Character, Stats, HpM, Items, Level)

	fmt.Println("\n", theGuy.Name, "\n", theGuy.Stats, "\n", theGuy.Life, "\n", theGuy.Items, "\n", theGuy.Level)

}
