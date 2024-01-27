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
    handleIputedComand(comand)
}
// Function to load the game introduction
func introduction() {
    name := " Super Powers Battle - Fantasy Game"
    version := 1.1
    fmt.Printf("\nWelcome to the %s!\n", name)
    fmt.Printf("Current version %v.\n", version)
    fmt.Printf("Your choice will unveil an unique experience!\n\n")
}
// Function to load the game Menu
func showMenu() {
    fmt.Println("Please choose a number:")
    fmt.Println("1- Start Game")
    fmt.Println("2- Load game")
    fmt.Printf("3- Exit\n")
}
// Function to read the input comand an return it
func readComand() int {
    var inputedComand int
    fmt.Scan(&inputedComand)
    fmt.Println("\nThe chosen comand was:", inputedComand)
    //fmt.Println("The variable memory address is:", &inputedComand)
    return inputedComand
}
// Function to handle the input comand (switch)
func handleIputedComand(comand int) {
    switch comand {
    case 1:
        fmt.Printf("\n################ Starting a new game...################\n")
        startGame()
    case 2:
        fmt.Println("Loading existing game..")
    case 3:
        fmt.Println("Exiting the program..")
        os.Exit(0)
    default:
        fmt.Println("Invalid choice. Please enter a number between 1 and 3.")
        os.Exit(-1)
    }
}
// Function to load the new game after user selection
func startGame() {
    var inputPlayer1 int
    var inputPlayer2 int
    var scorePlayer1 int
    var scorePlayer2 int
    fantasyLocations := []string{"Old Castle in the Mountains", "Elf's Hidden Realm", "Atlantis", "Asgard"}
    superPowers := []string{"Flying", "Invisibility", "Teleportation", "Time Manipulation"}
    
		//Using current time as a seed
    seed1 := time.Now().UnixNano()
    
		//User input - Player 1
    fmt.Println("Player 1: Please choose a number to start the game:")
    fmt.Scan(&inputPlayer1)
    
		//Modifying the seed with user input
    seed1 = (seed1 + int64(inputPlayer1)) % 100
    
		//Selecting the strings based on modified seed
    superPowerPlayer1 := superPowers[seed1%int64(len(superPowers))]
    fmt.Printf("Player 1 superpower is %s' .\n", superPowerPlayer1)
    //Using current time as a seed
    seed2 := time.Now().UnixNano()
    //User input - Player 2
    fmt.Println("Player 2: Please choose a number to start the game:")
    fmt.Scan(&inputPlayer2)
    
		//Modifying the seed with user input
    seed2 = (seed2 + int64(inputPlayer2)) % 100
    
		//Selecting the strings based on modified seed
    superPowerPlayer2 := superPowers[seed2%int64(len(superPowers))]
    fmt.Printf("Player 2 is %s' .\n", superPowerPlayer2)
    
		//Fantasy Location Randonly selected
    //Using current time as a seed
    seed := time.Now().UnixNano()
    seedSum := int64(inputPlayer1) + int64(inputPlayer2)
    seed = (seed + int64(seedSum)) % 100
    fantasyLocation := fantasyLocations[seed%int64(len(fantasyLocations))]
    fmt.Printf("Fantasy Location choosen is %s' .\n", fantasyLocation)
    
		//score of player1 based in the random selection of fantasy location and superpower.
    //In each location the superpowers has differents scores
    if fantasyLocation == "Old Castle in the Mountains" {
        if superPowerPlayer1 == "Flying" {
            scorePlayer1 = 4
        } else if superPowerPlayer1 == "Invisibility" {
            scorePlayer1 = 3
        } else if superPowerPlayer1 == "Teleportation" {
            scorePlayer1 = 2
        } else if superPowerPlayer1 == "Time Manipulation" {
            scorePlayer1 = 1
        } else {
            scorePlayer1 = 0
        }
    } else if fantasyLocation == "Elf's Hidden Realm" {
        if superPowerPlayer1 == "Invisibility" {
            scorePlayer1 = 4
        } else if superPowerPlayer1 == "Flying" {
            scorePlayer1 = 3
        } else if superPowerPlayer1 == "Time Manipulation" {
            scorePlayer1 = 2
        } else if superPowerPlayer1 == "Teleportation" {
            scorePlayer1 = 1
        }
    } else if fantasyLocation == "Atlantis" {
        if superPowerPlayer1 == "Teleportation" {
            scorePlayer1 = 4
        } else if superPowerPlayer1 == "Time Manipulation" {
            scorePlayer1 = 3
        } else if superPowerPlayer1 == "Invisibility" {
            scorePlayer1 = 2
        } else if superPowerPlayer1 == "Flying" {
            scorePlayer1 = 1
        }
    } else if fantasyLocation == "Asgard" {
        if superPowerPlayer1 == "Time Manipulation" {
            scorePlayer1 = 4
        } else if superPowerPlayer1 == "Teleportation" {
            scorePlayer1 = 3
        } else if superPowerPlayer1 == "Flying" {
            scorePlayer1 = 2
        } else if superPowerPlayer1 == "Invisibility" {
            scorePlayer1 = 1
        }
    }
    
		//score of player2 based in the random selection of fantasy location and superpower.
    //In each location the superpowers has differents scores
    if fantasyLocation == "Old Castle in the Mountains" {
        if superPowerPlayer2 == "Flying" {
            scorePlayer2 = 4
        } else if superPowerPlayer2 == "Invisibility" {
            scorePlayer2 = 3
        } else if superPowerPlayer2 == "Teleportation" {
            scorePlayer2 = 2
        } else if superPowerPlayer1 == "Time Manipulation" {
            scorePlayer2 = 1
        } else {
            scorePlayer2 = 0
        }
    } else if fantasyLocation == "Elf's Hidden Realm" {
        if superPowerPlayer2 == "Invisibility" {
            scorePlayer2 = 4
        } else if superPowerPlayer2 == "Flying" {
            scorePlayer2 = 3
        } else if superPowerPlayer2 == "Time Manipulation" {
            scorePlayer2 = 2
        } else if superPowerPlayer2 == "Teleportation" {
            scorePlayer2 = 1
        }
    } else if fantasyLocation == "Atlantis" {
        if superPowerPlayer2 == "Teleportation" {
            scorePlayer2 = 4
        } else if superPowerPlayer2 == "Time Manipulation" {
            scorePlayer2 = 3
        } else if superPowerPlayer2 == "Invisibility" {
            scorePlayer2 = 2
        } else if superPowerPlayer2 == "Flying" {
            scorePlayer2 = 1
        }
    } else if fantasyLocation == "Asgard" {
        if superPowerPlayer2 == "Time Manipulation" {
            scorePlayer2 = 4
        } else if superPowerPlayer2 == "Teleportation" {
            scorePlayer2 = 3
        } else if superPowerPlayer2 == "Flying" {
            scorePlayer2 = 2
        } else if superPowerPlayer2 == "Invisibility" {
            scorePlayer2 = 1
        }
    }
    //Compare the scores of the pçayers to see who won the battle
    if scorePlayer1 > scorePlayer2 {
        fmt.Printf("In %s, %s beats %s' - PLAYER 1 WINS!!!", fantasyLocation, superPowerPlayer1, superPowerPlayer2)
    } else if scorePlayer2 > scorePlayer1 {
        fmt.Printf("In %s, %s beats %s' - PLAYER 2 WINS!!!", fantasyLocation, superPowerPlayer2, superPowerPlayer1)
    } else {
        fmt.Println("It's a DRAW!!! Play again!!!")
    }
}