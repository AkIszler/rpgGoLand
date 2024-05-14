package main

import (
	"fmt"
	"rpgGo/rpgGo"
	"time"
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

	// fmt.Println(theGuy.Name, "\n", theGuy.Stats, "\n", theGuy.Life, "\n", theGuy.Items, "\n", theGuy.Level)

	fmt.Println("Well Hello", theGuy.Name, "\nWelcome to your story... you don't remember how you got here, \nbut your story is about to start\n\n ")
	time.Sleep(2 * time.Second)
	Start()

}

func Start() {

	fmt.Println(rpgGo.Intro.Story)
	time.Sleep(2 * time.Second)

	fmt.Println("1:Do you stay in the tavern\n2:Do you Leave the tavern?\n ")

	var ans string
	fmt.Scanln(&ans)

	if ans == "1" {
		fmt.Println(rpgGo.Stay.Story)
	}

}
