/** Author: Matthew Mckenney
 *  Date: 3/14/2020
 *  Purpose: create a guessing game
 *
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// gets the current date and time, as an integer
	seconds := time.Now().Unix()

	// seeds the random number generator
	rand.Seed(seconds)

	// generate an integer between 1 and 100
	target := rand.Intn(100) + 1

	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)

	success := false
	for guesses := 0; guesses < 10; guesses++ {

		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			success = true
			fmt.Println("Good job! you guessed it!")
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you did not guess my number. It was:", target)
	}
}
