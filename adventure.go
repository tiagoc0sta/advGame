package main

import (
	"fmt"
	"os"
	"time"
)

func main() {	

	introduction()
	showMenu()
	comand := readComand()

	switch comand {
		case 1:
			fmt.Println("Starting a new game..")
		case 2: 
			fmt.Println("Loading existing game..")
		case 3: 
			fmt.Println("Exiting the program..")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 3.")
			os.Exit(-1)

	}

	startGame()

		
	
	
 }

 func introduction(){
	name := "Fantasy Game"
	version := 1.3
	fmt.Printf("\nWelcome to the %s!\n",name)
	fmt.Printf("Current version %v.\n", version)
	fmt.Printf("Your choice will unveil an unique experience!\n\n")
 }

func showMenu() {
	fmt.Println("Please choose a number:")
	fmt.Println("1- Start Game")
	fmt.Println("2- Load game")
	fmt.Printf("3- Exit\n")
 }

 func readComand() int{
	var inputedComand int
	fmt.Scan(&inputedComand)
	
	fmt.Println("\nThe chosen comand was:", inputedComand)
	fmt.Println("The variable memory address is:", &inputedComand)

	return inputedComand
 }

 func startGame(){
	var input int
	

	fantasyLocations := []string{"Old Castle in the Mountains", "Elf's Hidden Realm", "Atlantis", "Asgard"}
	superPowers := []string{"Flying", "Invisibility", "Teleportation", "Time Manipulation"}

	//Using current time as a seed
	seed := time.Now().UnixNano()
	//User input
	fmt.Println("Please choose a number to start the game:")	
	fmt.Scan(&input)


	//Modifying the seed with user input
	seed = (seed + int64(input)) % 100

	//Selecting the strings based on modified seed
	choice1 := fantasyLocations[seed%int64(len(fantasyLocations))]
	choice2 := superPowers[seed%int64(len(superPowers))]

	fmt.Printf("Your fantansy is %s and your superpower is %s' .\n", choice1, choice2)

}