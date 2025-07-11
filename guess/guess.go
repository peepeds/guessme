package guess

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var (
	// normal mode
	easyLives   = 10
	mediumLives = 5
	hardLives   = 3

	// custom mode
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

func run(specialMode string) {
	target := randomize()

	var lives int

	if specialMode == "custom"{
		lives = customLives
	} else if specialMode == "challenge"{
		lives = math.MaxInt
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

func contains(slice []string, val string) bool {
    for _, item := range slice {
        if item == val {
            return true
        }
    }
    return false
}

func Play(selectedMode ...interface{}) {
	modes := []string{"normal", "custom", "challenge"}
	mode := "normal"

	// ========== 1. Handle CLI Flags ==========
	customLivesFlag := flag.Int("custom", -1, "Set custom lives for custom mode")
	challengeModeFlag := flag.Bool("challenge", false, "Enable challenge mode")
	flag.Parse()

	// ========== 2. Handle Variadic Args ==========
	if len(selectedMode) >= 1 {
		if m, ok := selectedMode[0].(string); ok && contains(modes, m) {
			mode = m
		}
	}

	// If mode is "custom", read the second arg
	if mode == "custom" && len(selectedMode) >= 2 {
		switch v := selectedMode[1].(type) {
		case int:
			setCustom(v)
		case string:
			if val, err := strconv.Atoi(v); err == nil {
				setCustom(val)
			} else {
				setCustom(easyLives)
				fmt.Println("Invalid customLives (string), using default:", easyLives)
			}
		default:
			setCustom(easyLives)
			fmt.Println("Unsupported customLives type, using default:", easyLives)
		}
	} else if mode == "custom" {
		// no value passed
		setCustom(easyLives)
		fmt.Printf("Custom mode requires number of lives, using default: %d\n", easyLives)
	}

	// ========== 3. Override Mode via Flags ==========
	if *customLivesFlag != -1 {
		mode = "custom"
		setCustom(*customLivesFlag)
	}

	if *challengeModeFlag {
		mode = "challenge"
	}

	// ========== 4. Start Game ==========
	welcome()

	for {
		run(mode)
		if !askReplay() {
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
