package main

import (
	"errors"
	"fmt"
)

func main() {

	printMe("Hey!")

	var numerator int = 11
	var denominator int = 5

	var result, remainder, err = intDivision(numerator, denominator)

	useIfTree := false

	// this is here so i don't have duplicated outputs
	if useIfTree {
		if err != nil {
			fmt.Printf(err.Error())
		} else if remainder == 0 {
			fmt.Printf("Division between %v and %v returns %v as a result.", numerator, denominator, result)
		} else {
			// using Printf to have better formatted strings
			fmt.Printf("Division between %v and %v returns %v as a result and %v as a remainder", numerator, denominator, result, remainder)
		}

	}

	// a more readable way of checking for errors
	// break is implied
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("Division between %v and %v returns %v as a result.", numerator, denominator, result)
	default:
		// using Printf to have better formatted strings
		fmt.Printf("Division between %v and %v returns %v as a result and %v as a remainder", numerator, denominator, result, remainder)
	}

}

// value string → I can't pass any other type as an argument
func printMe(value string) {
	fmt.Println(value)
}

// append the type at the end of declaration of a function
// to enforce what's going to be our return's type.

// using (int, int) we can specify that we want to return multiple values.
// return values types can be mixed.

// notice that this fuction can return an error, this is a common
// design pattern in Go.
func intDivision(numerator int, denominator int) (int, int, error) {

	var err error

	if denominator == 0 {
		err = errors.New("error: cannot divide by zero")
		return 0, 0, err
	}

	var divisionResult int = numerator / denominator
	var divisionRemainder int = numerator % denominator

	return divisionResult, divisionRemainder, err
}
