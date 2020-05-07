package main

import (
	"fmt"
	"math"
	"time"
)

func compoundInterest() {
	// A = P(1+r/n)^(nt)
	// A = the future value of the investment/loan, including interest
	// P = the principal investment amount (the initial deposit or loan amount)
	// r = the annual interest rate (decimal)
	// n = the number of times that interest is compounded per unit t
	// t = the time the money is invested or borrowed for

	var a, p, r, n, t float64

	// user enters principal value and reads into variable
	fmt.Print("Please enter the principal value: $")
	_, _ = fmt.Scanln(&p)

	// user enters the annual interest rate and reads into variable
	fmt.Print("Please enter the annual interest rate: ")
	_, _ = fmt.Scanln(&r)

	//r /= 100

	// user enters the number of times that interest is compounded per unit t and reads into variable
	fmt.Print("Please enter the number of times that interest is compounded per unit t: ")
	_, _ = fmt.Scanln(&n)

	// user enters the time in years and reads into variable
	fmt.Print("Please enter the amount of time in years: ")
	_, _ = fmt.Scanln(&t)

	// A = P(1+r/n)^(nt)
	a = p * math.Pow(1+(r/n), n*t)

	fmt.Printf("the future value of your investment/loan of $%.2f is $%.2f in %.0f years.\n\n", p, a, t)
}

func simpleInterest() {

	// simple interest calculation
	// p = principal
	// r = rate
	// t = time in years
	// I = interest rate
	var p, r, t float64

	// user enter p = principal value
	fmt.Print("Please enter the principle value: $")
	_, _ = fmt.Scanln(&p)

	// user enter r = rate
	fmt.Print("Please enter the rate as a whole number: ")
	_, _ = fmt.Scanln(&r)

	// turns the user whole number into a percentage
	r /= 100

	// user enter t = time
	fmt.Print("Please enter the amount of years: ")
	_, _ = fmt.Scanln(&t)

	// simple interest formula calculation
	i := p * r * t

	// print out simple interest
	fmt.Printf("$%.2f\n\n", i)
}

func welcome() {
	fmt.Printf("Welcome to my program\n")
	fmt.Printf(time.Now().Format("01-02-2006\n"))

}

func main() {
	welcome()
	var drive int
	fmt.Print("please enter 1 for simple interest calculator or 2 for compound: ")
	_, _ = fmt.Scanln(&drive)

	switch drive {
	case 1:
		simpleInterest()
	case 2:
		compoundInterest()
	default:
		fmt.Printf("%d is not a valid input!", drive)

	}

}
