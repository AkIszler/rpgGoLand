package main

import (
	"fmt"
	"rpgGo/rpgGo"
)

func pickClass(class string) rpgGo.Class {
	var classPicked rpgGo.Class
	switch class {
	case "Wizard":
		classPicked = rpgGo.WizardClass
	case "Warrior":
		classPicked = rpgGo.WarriorClass
	default:
		classPicked = rpgGo.RogueClass
	}

	return classPicked
}

func BuildCharacter(cn string, class rpgGo.Class) ([]string, map[string]int) {
	characterStats := make(map[string]int)
	var character []string

	character = append(character, cn)
	character = append(character, class.Name)

	characterStats["HP"] = class.HPchange
	characterStats["Atk"] = class.Atkchange
	characterStats["Def"] = class.Defchange
	characterStats["Int"] = class.Intchange
	characterStats["Dex"] = class.Dexchange
	characterStats["Wis"] = class.Wischange
	characterStats["Cha"] = class.Chachange

	return character, characterStats
}

func main() {

	fmt.Println("Welcome to GoLang Island!")

	fmt.Println("What is your name?")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)
	fmt.Println("What class are you?")
	var class string
	fmt.Scanln(&class)

	classPick := pickClass(class)

	fmt.Println(BuildCharacter(name, classPick))

}
