package rpgGo

type Multiplier struct {
	Base       int
	Multiplier int
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
	Dmg:          2,
	RequireMagic: false,
}

var Bow = Items{
	Name:         "Bow",
	ID:           4,
	Dmg:          1 * 3,
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
