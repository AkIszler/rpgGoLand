package rpgGo

import (
	"fmt"
)

type Character struct {
	Name  string
	Class string
	Level int
	HP    int
	Mana  int
	Atk   int
	Def   int
	Int   int
	Dex   int
	Wis   int
	Cha   int

	// Add more fields as needed
}

var baseCharacter = Character{
	Name:  "",
	Class: "",
	Level: 1,
	HP:    10,
	Mana:  10,
	Atk:   4,
	Def:   4,
	Int:   4,
	Dex:   4,
	Wis:   4,
	Cha:   4,

	// Add more fields as needed
}

type Class struct {
	Name       string
	HPchange   int
	ManaChange int
	Atkchange  int
	Defchange  int
	Intchange  int
	Dexchange  int
	Wischange  int
	Chachange  int
}

var WizardClass = Class{

	Name:       "Wizard",
	HPchange:   baseCharacter.HP + 3,
	ManaChange: baseCharacter.Mana + 6,
	Atkchange:  baseCharacter.Atk + 3,
	Defchange:  baseCharacter.Def + 3,
	Intchange:  baseCharacter.Int - 2,
	Dexchange:  baseCharacter.Dex - 1,
	Wischange:  baseCharacter.Wis - 2,
	Chachange:  baseCharacter.Cha - 2,
}

var WarriorClass = Class{
	Name:       "Warrior",
	HPchange:   baseCharacter.HP + 3,
	ManaChange: baseCharacter.Mana - 3,
	Atkchange:  baseCharacter.Atk + 3,
	Defchange:  baseCharacter.Def + 3,
	Intchange:  baseCharacter.Int - 2,
	Dexchange:  baseCharacter.Dex - 1,
	Wischange:  baseCharacter.Wis - 2,
	Chachange:  baseCharacter.Cha - 2,
}

var RogueClass = Class{
	Name:       "Rogue",
	HPchange:   baseCharacter.HP,
	ManaChange: baseCharacter.Mana - 2,
	Atkchange:  baseCharacter.Atk + 1,
	Defchange:  baseCharacter.Def - 2,
	Intchange:  baseCharacter.Int,
	Dexchange:  baseCharacter.Dex + 5,
	Wischange:  baseCharacter.Wis + 2,
	Chachange:  baseCharacter.Cha + 1,
}

// Define some sample data

func PickClass() Class {
	var classPicked Class
	classpicked := ""

	for {
		fmt.Println("What class are you?")
		fmt.Println("1. Wizard\n2. Warrior\n3. Rogue")
		fmt.Scanln(&classpicked)

		if classpicked == "1" {
			classPicked = WizardClass
			break
		} else if classpicked == "2" {
			classPicked = WarriorClass
			break
		} else if classpicked == "3" {
			classPicked = RogueClass
			break
		} else {
			fmt.Println("Invalid input")
		}
	}

	return classPicked
}

func BuildCharacter(cn string, class Class) ([]string, map[string]int, map[string]int, map[string]int) {
	characterStats := make(map[string]int)
	characterHpMana := make(map[string]int)
	var character []string
	var items = GetStartingItems(class)
	//name
	character = append(character, cn)
	character = append(character, class.Name)
	//stats
	characterStats["Atk"] = class.Atkchange
	characterStats["Def"] = class.Defchange
	characterStats["Int"] = class.Intchange
	characterStats["Dex"] = class.Dexchange
	characterStats["Wis"] = class.Wischange
	characterStats["Cha"] = class.Chachange
	//hp Manas
	characterHpMana["HP"] = class.HPchange
	characterHpMana["Mana"] = class.ManaChange
	//get base gear
	return character, characterStats, characterHpMana, items
}

type Items struct {
	Name         string
	ID           int
	Dmg          int
	RequireMagic bool
}

var Sword = Items{
	Name:         "Sword",
	ID:           1,
	Dmg:          5,
	RequireMagic: false,
}

var Staff = Items{
	Name:         "Staff",
	ID:           2,
	Dmg:          2,
	RequireMagic: true,
}

var Knife = Items{
	Name:         "Knife",
	ID:           3,
	Dmg:          1 * 2,
	RequireMagic: false,
}

func GetStartingItems(class Class) map[string]int {

	var item = make(map[string]int)

	if class == WarriorClass {
		item[Sword.Name] = Sword.Dmg
	} else if class == WizardClass {
		item[Staff.Name] = Staff.Dmg
	} else if class == RogueClass {
		item[Knife.Name] = Knife.Dmg
	}
	return item
}
