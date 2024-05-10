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

	character := rpgGo.Character{
		Name:  cn,
		Level: 1,
		Class: classPicked,
		HP:    base.HP + classInfo.HPchange,
		Atk:   base.Atk + classInfo.Atkchange,
		Def:   base.Def + classInfo.Defchange,
		Int:   base.Int + classInfo.Intchange,
		Wis:   base.Wis + classInfo.Wischange,
		Cha:   base.Cha + classInfo.Chachange,
	}
	fmt.Println(character)
}

func main() {

	base := rpgGo.Character{}
	BuildCharacter("Tony", "Warrior", base)
	fmt.Println("Hello, 世界")

}
