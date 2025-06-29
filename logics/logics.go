package logics

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func Init(){
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")

	play()
}

func modes()int {
	fmt.Println("\nPlease Choose game difficulty level")
	fmt.Println("1. Easy (10 choices)")
	fmt.Println("2. Medium (5 choices)")
	fmt.Println("3. Hard (3 choices)")
	fmt.Print(">> ")
	var mode int
	fmt.Scan(&mode)
	return mode
}

func randomize() int64 {
	random, err := rand.Int(rand.Reader, big.NewInt(100)) // range: 0 - 99
	if err != nil {
		fmt.Println("Error generating random number:", err)
		return 0
	}
	return random.Int64() + 1
}

func choices() int{
	var choice int
	fmt.Print("\nEnter your choices : ")
	fmt.Scan(&choice)
	return choice
} 


func play() (result, tries int) {
	var lives int 
	random := int(randomize())
	running := true
	mode := modes()

	tries = 0 
	result = 0
	switch{
		case mode == 1 :
			lives = 10
		case mode == 2 : 
			lives = 5
		case mode == 3 : 
			lives = 3
	} 
	for running {
		choise := choices()
		clues(choise, random)
		tries++

		if	tries == lives {
			running = false
		}

		if choise == random {
			result = 1
			running = false
		} 
		
	}

	fmt.Println(result,tries)
	results(result,tries)
	return result, tries
}

func clues(choice, target int){
	if choice < target {
		fmt.Println("To Low")
	} else if choice > target {
		fmt.Println("To Big")
	} else {
		fmt.Println("Win!")
	}
}

func results(result, tries int){
	if result == 1 {
		fmt.Printf("Winning with : %d tries\n",tries)
	} else {
		fmt.Printf("Better try again :v\n")
	}
	
}

