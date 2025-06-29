package guess

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
)

var (
	easyLives   = 10
	mediumLives = 5
	hardLives   = 3
	customLives int
)

func SetUp(easy, medium, hard int){
	easyLives = easy
	mediumLives = medium
	hardLives = hard
}

func setCustom(custom int){
	customLives = custom
}

func welcome() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
}

func selectMode() int {
	fmt.Println("\nPlease choose game difficulty level:")
	fmt.Println("1. Easy (10 lives)")
	fmt.Println("2. Medium (5 lives)")
	fmt.Println("3. Hard (3 lives)")
	fmt.Print(">> ")

	var mode int
	fmt.Scan(&mode)

	switch mode {
	case 1:
		return easyLives
	case 2:
		return mediumLives
	case 3:
		return hardLives
	default:
		fmt.Println("Invalid mode. Defaulting to Easy.")
		return easyLives
	}
}

func randomize() int {
	random, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		fmt.Println("Error generating random number:", err)
		return 1
	}
	return int(random.Int64()) + 1
}

func getGuess() int {
	var guess int
	fmt.Print("Enter your guess: ")
	fmt.Scan(&guess)
	return guess
}

func giveClue(guess, target int) {
	if guess < target {
		fmt.Println("Too low.")
	} else if guess > target {
		fmt.Println("Too high.")
	} else {
		fmt.Println("Correct!")
	}
}

func run(isCustom bool) {
	target := randomize()
	var lives int
	if isCustom{
		lives = customLives
	} else {
		lives = selectMode()
	}

	for attempts := 1; attempts <= lives; attempts++ {
		guess := getGuess()
		giveClue(guess, target)

		if guess == target {
			fmt.Printf("🎉 You won in %d tries!\n", attempts)
			return
		} else if attempts == lives {
			fmt.Printf("💥 You lost. The number was %d.\n", target)
		}
	}
}

func Play(){
	isCustomMode := false

	customLives := flag.Int("custom",-1, "Set Custom lives / playground modes") 
	flag.Parse()

	if *customLives != -1 {
		isCustomMode = true
		setCustom(*customLives)
		welcome()
	}else {
		welcome()
	}

	for {
		run(isCustomMode)
		if !askReplay(){
			fmt.Println("Thanks for playing")
			break
		}
	}
}

func askReplay() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Play again? [y/n]: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))
	return input == "y"
}
