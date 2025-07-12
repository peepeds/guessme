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
	// CLI flags
	CustomLivesFlag   = flag.Int("custom", -1, "Set custom lives for custom mode")
	ChallengeModeFlag = flag.Bool("challenge", false, "Enable challenge mode")

	// Default lives per mode
	easyLives   = 10
	mediumLives = 5
	hardLives   = 3
	customLives int

	// Modes
	listModes = []string{"normal", "custom", "challenge"}

	// flag state tracker
	flagParsed = false
)

func ensureFlagParsed() {
	if !flagParsed {
		flag.Parse()
		flagParsed = true
	}
}

// === Setup lives ===
func SetUp(easy, medium, hard int) {
	easyLives = easy
	mediumLives = medium
	hardLives = hard
}

// === Game setup ===

func setLives(selectedMode string) int {
	switch selectedMode {
	case "custom":
		if customLives == 0 {
			return easyLives
		}
		return customLives
	case "challenge":
		return math.MaxInt
	default:
		return selectMode()
	}
}

func validatesMode(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func setMode(modes ...interface{}) string {
	ensureFlagParsed()

	var mode string

	if len(modes) == 0 {
		mode = "normal"
	} else if m, ok := modes[0].(string); ok && validatesMode(listModes, m) {
		mode = m
	}

	if mode == "custom" && len(modes) >= 2 {
		switch v := modes[1].(type) {
		case int:
			customLives = v
		case string:
			if val, err := strconv.Atoi(v); err == nil {
				customLives = val
			} else {
				customLives = easyLives
			}
		default:
			customLives = easyLives
		}
	} else if mode == "custom" && len(modes) == 1 { // if user not defined their lives
		customLives = easyLives
	}

	if *CustomLivesFlag != -1 {
		mode = "custom"
		customLives = *CustomLivesFlag
	}
	if *ChallengeModeFlag {
		mode = "challenge"
	}

	return mode
}

func welcome() {
	fmt.Println("🎮 Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
}

func selectMode() int {
	fmt.Println("\nChoose difficulty:")
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
		fmt.Println("Invalid input. Defaulting to Easy.")
		return easyLives
	}
}

func randomize() int {
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return 1
	}
	return int(n.Int64()) + 1
}

// === Game Loop ===

func getGuess() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your guess: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if guess, err := strconv.Atoi(input); err == nil {
			return guess
		}
		fmt.Println("❌ Invalid input. Please enter a number.")
	}
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

func startGuessingRound(target int, lives int) int {
	for attempts := 1; attempts <= lives; attempts++ {
		guess := getGuess()
		giveClue(guess, target)

		if guess == target {
			fmt.Printf("🎉 You won in %d tries!\n", attempts)
			return attempts
		} else if attempts == lives {
			fmt.Printf("💥 You lost. The number was %d.\n", target)
		}
	}
	return -1
}

func run(mode string) int {
	target := randomize()
	lives := setLives(mode)
	return startGuessingRound(target, lives)
}

func askReplay() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Play again? [y/n]: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))
	return input == "y"
}

// === Entrypoint ===

func Play(modes ...interface{}) {
	mode := setMode(modes...)
	welcome()

	for {
		run(mode)
		if !askReplay() {
			fmt.Println("👋 Thanks for playing!")
			break
		}
	}
}
