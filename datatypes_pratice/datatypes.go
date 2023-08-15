package main

import "fmt"

// Grouped constants, x is untyped, y is typed
const (
	x                     = 10
	y               int32 = 15
	applicationName       = "Lesson 7"
	isRunning             = false
	isTrue                = 1 < 2
)

const (
	Zero int = iota
	One
	Two
	Three
	Four
	Five
)

func main() {
	var smallPositiveValue int8
	smallPositiveValue = 10
	fmt.Println(smallPositiveValue)

	var myInt int = 2083234
	fmt.Println(myInt)

	// byte is mainly used to represent raw data and also to
	// distinguish between uint8 and byte since byte is an
	// alias for uint8
	var myByte byte = 'A'
	fmt.Println(myByte)

	// Go uses UTF-8 to encode Unicode
	var myRune rune = '💫'
	fmt.Println(myRune)

	var smallFloat float32
	fmt.Printf("Floats initialize as 0: %v\n", smallFloat)
	smallFloat = 23.023494333
	fmt.Printf("As you can see, precision is limited with float32: %v\n", smallFloat)

	// Default float type is float64
	var bigFloat float64
	bigFloat = 23.29350289358092385234999999999
	fmt.Printf("Precision is better, but still limited with float64: %v\n", bigFloat)

	// Default complex type is complex128
	var myComplex complex128
	var realBigFloat float64 = 99.2348923589135813511111
	var imagBigFloat float64 = 38.109831581024892948248888

	myComplex = complex(realBigFloat, imagBigFloat)
	fmt.Printf("Complex number with real part & imaginary part concatenated by +: %v\n", myComplex)

	// You can declare 2 variables of the same type on 1 line
	var myRealPart, myImaginaryPart float64
	myRealPart = real(myComplex)
	myImaginaryPart = imag(myComplex)
	fmt.Printf("Real part: %v, Imaginary part: %v\n", myRealPart, myImaginaryPart)

	// Strings are immutable (you can't change them directly),
	// but you can update the variable holding the strong. They
	// are stored as a byte array internally. UTF-8 encoding: each
	// Unicode character can take from 1 to 4 bytes in memory.
	var test string = `Raw strings let you do this:
New-line
	Tab`
	fmt.Printf(test)
	first := "Trevor"
	last := "English"
	fmt.Printf("\nMy name is %s %s", first, last)

	fullName := fmt.Sprintf("\nMy name is %s %s", first, last)
	fmt.Println(fullName)

	age, name := 10, "Code & learn"
	fmt.Printf("You can declare 2 variables in one line: %d, %s\n", age, name)

	var a int = x
	fmt.Println(a)

	// We can assign the const x to variable b because it is untyped
	var b float64 = x
	fmt.Println(b)

	// For typed constants, you need to cast them to whatever the type of the variable is
	var c int = int(y)
	fmt.Println(c)

	// We can't have a runtime constant
	const z = complex(10.2, 100.9)

	// Demonstrating how Iota increments integers:
	fmt.Printf("Iota increment: %v, %v, %v, %v, %v, %v\n", Zero, One, Two, Three, Four, Five)

}
