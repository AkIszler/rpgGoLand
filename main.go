package main

import (
	"fmt"
	"rpgGo/rpgGo"
)

func pickClass(class string) (string, rpgGo.Class) {
	var classPicked string
	var classPickedClass rpgGo.Class
	switch class {
	case "Wizard":
		classPicked = rpgGo.WizardClass.Name
	case "Warrior":
		classPicked = rpgGo.WarriorClass.Name
	default:
		classPicked = rpgGo.RogueClass.Name
	}

	return classPicked, classPickedClass
}

func BuildCharacter(cn string, class string, base rpgGo.Character) {
	classPicked, classInfo := pickClass(class)
	fmt.Println(classPicked)
	fmt.Println(base)
	character := rpgGo.Character{
		Name:  cn,
		Level: 1,
		Class: classPicked,
		HP:    base.HP + classInfo.HPchange,
		Atk:   base.Atk + classInfo.Atkchange,
		Def:   base.Def + classInfo.Defchange,
		Int:   base.Int + classInfo.Intchange,
		Dex:   base.Dex + classInfo.Dexchange,
		Wis:   base.Wis + classInfo.Wischange,
		Cha:   base.Cha + classInfo.Chachange,
	}
	fmt.Println(character)
}

func main() {

	base := rpgGo.Character{
		Name:  "",
		Class: "",
		Level: 1,
		HP:    10,
		Atk:   4,
		Def:   4,
		Int:   4,
		Dex:   4,
		Wis:   4,
		Cha:   4,
		// Add more fields as needed
	}
	BuildCharacter("Tony", "Warrior", base)
	fmt.Println("Hello, 世界")

}
