package main

import "fmt"

func main() {
	var intArr [3]int32 // length cannot change after inizialization.

	intArr[2] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[0:3]) // we are printing from index 0 to index 2.

	// we are printing the memory address of a specified index
	fmt.Println(&intArr[0]) // 0xc000096080
	fmt.Println(&intArr[1]) // 0xc000096084
	fmt.Println(&intArr[2]) // 0xc000096088

	// notice that arrays store info in contiguous memory locations.
	// This is really beneficial for the compiler because it just needs to know where the first byte is
	// and then ncrement the address by N bytes to access the next index (4 bytes in this case)

	// we can immediately initialize this array
	var arr [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(arr)

	// again, type can be inferred, notice that size can be inferred too by
	// using the ... syntax
	inferredArray := [...]int32{1, 2, 3, 4, 5, 6}
	fmt.Println(inferredArray)

	// Slices are just array wrappers with extended functionalities
	intSlice := []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("Slice length: %v - Capacity: %v \n", len(intSlice), cap(intSlice))

	// unlike arrays, slices can change size
	intSlice = append(intSlice, 7) // use append() to add an elementat the end of an array
	fmt.Println(intSlice)
	fmt.Printf("Slice length: %v - Capacity: %v \n", len(intSlice), cap(intSlice))

	// when using append, we might not actually extend the size of our pre-existing slice,
	// but instead, a new one will be given to us with a higher capacity.

	// this is because in Go a slice can have a certain capacity, which basically means the
	// number of values that can be stored.

	// the append function first looks at the capacity of our array, if it's full, it will
	// create a new one appending our value with a higher capacity.

	// we can create slices using the make() function, we can optionally pass
	// the capacity as an argument if we roughly know it.
	// this avoids having to re-allocate the array.

	var intSlice2 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice2)
	fmt.Printf("Slice length: %v - Capacity: %v \n", len(intSlice2), cap(intSlice2))

	mapsFunction()
}

func mapsFunction() {

	fmt.Println("")

	// maps are sets of key-value pairs
	// we can look for a specific value by accessing its key.

	// map declaration
	// 	- keys are of type string
	// 	- values are an 8-bit integer
	var map1 map[string]uint8 = make(map[string]uint8)
	fmt.Println(map1)

	var map2 = map[string]uint8{"Henri": 20, "Kevin": 29}
	fmt.Println(map2)
	fmt.Printf("Henri is %v \n", map2["Henri"])

	// notice that maps return the default value for the type specified when accessing
	// a key that does not exist.
	fmt.Println(map2["John"])

	// maps can also return an optional boolean value to differentiate
	// when the key exists.

	// this means we can do something like this.
	var age, exists = map2["John"]

	if exists {
		fmt.Printf("The age is %v \n", age)
	} else {
		fmt.Println("Invalid name.")
	}

	// we can delete values from maps by using the delete() function
	// as deletions occur by reference, no return value are given.
	delete(map2, "Kevin")

	// we can iterate over arrays, slices and maps like this
	// when iterating over a map, no order is preserved
	// specifying another variable name after "key", will
	// give us the value
	for key, value := range map2 {
		fmt.Printf("Name: %v - Age: %v \n", key, value)
	}

	// the same principle can be used for arrays and slices
	// but we will have index and value instead.

	// we can do "fixed length runs" by using this syntax
	idx := 0
	for idx < 10 {
		fmt.Println(idx)
		idx++
	}

	// we can forget the condition and break our loop manually.
	jdx := 0
	for {
		if jdx >= 10 {
			break
		}
		fmt.Println(jdx)
		jdx++
	}

	// the most concise way of declaring a for loop is this, though
	for fdx := 0; fdx < 11; fdx++ {
		fmt.Println(fdx)
	}
}
