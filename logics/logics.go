package logics

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
)

const (
	easyLives   = 10
	mediumLives = 5
	hardLives   = 3
)

func Init() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")

	play()
	fmt.Println("Thanks for playing!")
}

// Menampilkan pilihan mode dan mengembalikan jumlah nyawa
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

// Menghasilkan angka acak dari 1 hingga 100
func randomize() int {
	random, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		fmt.Println("Error generating random number:", err)
		return 1
	}
	return int(random.Int64()) + 1
}

// Meminta input angka dari user
func getGuess() int {
	var guess int
	fmt.Print("Enter your guess: ")
	fmt.Scan(&guess)
	return guess
}

// Memberikan petunjuk berdasarkan tebakan user
func giveClue(guess, target int) {
	if guess < target {
		fmt.Println("Too low.")
	} else if guess > target {
		fmt.Println("Too high.")
	} else {
		fmt.Println("Correct!")
	}
}

// Memainkan satu ronde permainan
func play() {
	target := randomize()
	lives := selectMode()

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

// Menanyakan apakah user ingin bermain lagi
func askReplay() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Play again? [y/n]: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))
	return input == "x"
}
