package rpgGo

type Character struct {
	Name  string
	Class string
	Level int
	HP    int
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
	Atk:   4,
	Def:   4,
	Int:   4,
	Dex:   4,
	Wis:   4,
	Cha:   4,

	// Add more fields as needed
}

type Class struct {
	Name      string
	HPchange  int
	Atkchange int
	Defchange int
	Intchange int
	Dexchange int
	Wischange int
	Chachange int
}

var WizardClass = Class{
	Name:      "Wizard",
	HPchange:  2,
	Atkchange: -1,
	Defchange: -2,
	Intchange: +5,
	Dexchange: +1,
	Wischange: +2,
	Chachange: +0,
}

var WarriorClass = Class{
	Name:      "Warrior",
	HPchange:  6,
	Atkchange: +5,
	Defchange: +3,
	Intchange: -1,
	Dexchange: 0,
	Wischange: -1,
	Chachange: -1,
}

var RogueClass = Class{
	Name:      "Rogue",
	HPchange:  0,
	Atkchange: +1,
	Defchange: -2,
	Intchange: -1,
	Dexchange: +4,
	Wischange: +1,
	Chachange: +3,
}

// Define some sample data
