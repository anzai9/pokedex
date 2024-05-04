package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex: ")
	fmt.Println("Usage: <command> <arguments>")
	fmt.Println()
	fmt.Println("Commands:")
	for _, cmd := range commandMap() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}