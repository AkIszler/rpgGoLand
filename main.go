package main

import (
	"fmt"
	"rpgGo/rpgGo"
)

func main() {

	fmt.Println("Hello, 世界")

	for _, character := range rpgGo.Characters {
		fmt.Printf("Name: %s, Class: %s, Level: %d, HP: %d\n", character.Name, character.Class, character.Level, character.HP)
	}
}
