package main

import (
	"fmt"
	"rpgGo/rpgGo"
)

// func pickClass() rpgGo.Class {
// 	var classPicked rpgGo.Class
// 	classpicked := ""

// 	for {
// 		fmt.Println("What class are you?")
// 		fmt.Println("1. Wizard\n2.Warrior\n3.Rogue")
// 		fmt.Scanln(&classpicked)

// 		if classpicked == "1" {
// 			classPicked = rpgGo.WizardClass
// 			break
// 		} else if classpicked == "2" {
// 			classPicked = rpgGo.WarriorClass
// 			break
// 		} else if classpicked == "3" {
// 			classPicked = rpgGo.RogueClass
// 			break
// 		} else {
// 			fmt.Println("Invalid input")
// 		}
// 	}

// 	return classPicked
// }

// func BuildCharacter(cn string, class rpgGo.Class) ([]string, map[string]int, map[string]int) {
// 	characterStats := make(map[string]int)
// 	characterHpMana := make(map[string]int)
// 	var character []string

// 	character = append(character, cn)
// 	character = append(character, class.Name)

// 	characterStats["Atk"] = class.Atkchange
// 	characterStats["Def"] = class.Defchange
// 	characterStats["Int"] = class.Intchange
// 	characterStats["Dex"] = class.Dexchange
// 	characterStats["Wis"] = class.Wischange
// 	characterStats["Cha"] = class.Chachange

// 	characterHpMana["HP"] = class.HPchange
// 	characterHpMana["HP"] = class.ManaChange

// 	return character, characterStats, characterHpMana
// }

func main() {

	fmt.Println("Welcome to GoLang Island!")

	fmt.Println("What is your name?")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)

	classPick := rpgGo.PickClass()

	fmt.Println(rpgGo.BuildCharacter(name, classPick))

}
