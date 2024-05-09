package rpgGo

type Character struct {
	Name  string
	Class string
	Level int
	HP    int
	Atk   int
	Def   int
	Int   int
	Wis   int
	Cha   int

	// Add more fields as needed
}

// Define some sample data
var Characters = []Character{
	{"Alice", "open", 10, 100, 4, 4, 4, 4, 4},
	// Add more characters as needed
}
