package main

import "fmt"

func main() {
	printMe("A function printed this using an argument!")
}

// value string → I can't pass any other type as an argument
func printMe(value string) {
	fmt.Println(value)

	fmt.Println(intDivision(5, 2))

}

// append the type at the end of declaration of a function
// to enforce what our function returns
func intDivision(numerator int, denominator int) int {
	return numerator / denominator
}
