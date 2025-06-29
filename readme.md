# 🎮 Guess Me

A simple CLI-based number guessing game written in Go.  
Guess the random number with multiple difficulty modes or set your own custom mode!

## ✨ Features

- 🔁 Multiple rounds gameplay  
- ⚙️ Set custom modes with user-defined lives  
- 🧠 Default difficulty modes: Easy, Medium, Hard  
- 🧪 CLI-based flag inputs  

## 🛠️ Prerequisites

- **Go v1.24** or higher  
Check if Go is installed:
	```bash
	go version
	```

##  🚀 Installation
1. Initialize your Go project
	```bash
	go mod init your-project-name
	```
2.  Install Guess Me package
	```bash
	go get github.com/peepeds/guessme@latest
	```
3. Create the entry file
	```bash
	touch main.go
	```

##  🧪 Usage
### ✅ Skeleton Code
```go
package main

import "github.com/peepeds/guessme/guess"

func main() {
    // Setup default lives for modes: easy, medium, hard
    //optional
    guess.SetUp(25, 10, 5)

    // Start the game
    guess.Play()
}

```

### ▶️ Run the game
```go
go run main.go
```

### 🧩 With custom mode flag
``` go
go run main.go --custom 10
```
