package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(100) + 1

	var guess int
	attempts := 0

	fmt.Println("Guess a number between 1 and 100")

	for {
		fmt.Print("Enter guess: ")
		fmt.Scan(&guess)
		attempts++

		if guess < number {
			fmt.Println("Too low")
		} else if guess > number {
			fmt.Println("Too high")
		} else {
			fmt.Println("Correct!")
			fmt.Println("Attempts:", attempts)
			break
		}
	}
}

