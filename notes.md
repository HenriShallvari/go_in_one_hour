# Learning GoLang as fast as I can


## Facts about Go
1. Go is strongly typed
2. Types can be explicitly written or inferred.
3. Go offers greater performance against interpreted languages such as Python.


## Introduction
Golang files **need** to have at the top a package declaration line.

`package main` is actually a special package name, as this tells the compiler that this file is the entry point of our program.

As the main package is quite special, we must have a main function.

Functions are declared like this:

```
func main() {
	// some code
}
```

## Constants, variables and basic datatypes

We have various types of integers, such as:
- **int** → classic integer
- **int16** → 16-bit integer
- **int32** → 32-bit integer
- **int64** → 64-bit integer

The maximum value that we can store in a 16-bit integer is **32767**, this means that if we do something like `var num int16 = 32767 + 1` we'll get an overflow error.

BUT

If we have something like
```
var num int16 = 32767
num += 1
fmt.Println(num)
```
we won't get neither a compilation error nor a runtime error, our program will simply output **-32768**.

We also have access to unsigned integers (`uint`), which are available in the same bit sizes as integers but can store only positive values.

We also have floats, which in Go need to be "sized". 

To clarify, this means that we can only choose from either `float32` or `float64`.

These are actually not that precise and the degree of precision that we'll get from values stored in **float variables** depends on how that number is actually stored in memory.

As Go is strongly typed, it isn't actually possible to do cross-type arithmetic operations.

This throws an exception:

```
var intNum32 int32 = 2
var floatNum32 float32 = 10.1
var result float32 = floatNum32 + intNum32
```

We can, instead, cast our integer like this:

```
var result float32 = floatNum32 + float32(intNum32)
```

Strings are used to store... strings!
We can declare them like this:

```
var myString string = "Hello World"
```

We can even concatenate them:
```
var myString string = "Hello"+ " " + "World"
```

We can retrieve the length of a string by using the `len` function:
```
len(myString)
```
**This does not return the number of characters but the number of bytes.**

As Go uses UTF-8 encoding, characters outside the 256 ASCII character set use more than one byte.

To get the length of a string which uses "special" characters, we can import the built-in **package unicode/utf8.** and use the function `RuneCountInString()`, like this:
```
package main

import "fmt"
import "unicode/utf8"

func main(){
    // some code

    fmt.Println(utf8.RuneCountInString("STRING_WITH_SPECIAL_CHARS"))
}
```

**Runes are another type used to store strings.**

To declare constants, just use const instead of var.