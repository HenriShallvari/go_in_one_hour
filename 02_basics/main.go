package main

import "fmt"

func main() {
	var intNum int16 = 32767
	intNum += 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32) // we need to cast our integer to sum it
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2) // normal division
	fmt.Println(intNum1 % intNum2) // gets the remainder of the division

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	var boolVal bool = true
	fmt.Println(boolVal)

	// other types of variable declarations:
	var myInt int    // Go assigns 0 as the default value
	var myStr string // Go assigns '' as the default value

	var myVar = "text" // type is inferred
	myVar2 := "text"   // type is inferred, I prefer this style

	fmt.Println(myInt)
	fmt.Println(myStr)
	fmt.Println(myVar)
	fmt.Println(myVar2)

	//example of declaring multiple variables at once
	var1, var2 := 1, 2
	fmt.Println(var1, var2)
}
