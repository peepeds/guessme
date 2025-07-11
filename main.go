package main

import "github.com/peepeds/guessme/guess"

func main(){
	guess.SetUp(10,8,5)
	guess.Play("custom", 90)
}