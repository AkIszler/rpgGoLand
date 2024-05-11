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
	HPchange:  baseCharacter.HP + 3,
	Atkchange: baseCharacter.Atk + 3,
	Defchange: baseCharacter.Def + 3,
	Intchange: baseCharacter.Int - 2,
	Dexchange: baseCharacter.Dex - 1,
	Wischange: baseCharacter.Wis - 2,
	Chachange: baseCharacter.Cha - 2,
}

var WarriorClass = Class{
	Name:      "Warrior",
	HPchange:  baseCharacter.HP + 3,
	Atkchange: baseCharacter.Atk + 3,
	Defchange: baseCharacter.Def + 3,
	Intchange: baseCharacter.Int - 2,
	Dexchange: baseCharacter.Dex - 1,
	Wischange: baseCharacter.Wis - 2,
	Chachange: baseCharacter.Cha - 2,
}

var RogueClass = Class{
	Name:      "Rogue",
	HPchange:  baseCharacter.HP + 3,
	Atkchange: baseCharacter.Atk + 3,
	Defchange: baseCharacter.Def + 3,
	Intchange: baseCharacter.Int - 2,
	Dexchange: baseCharacter.Dex - 1,
	Wischange: baseCharacter.Wis - 2,
	Chachange: baseCharacter.Cha - 2,
}

// Define some sample data
