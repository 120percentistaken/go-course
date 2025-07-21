package main

import "fmt"

//constants are immutable values that cannot be changed once declared.
func main() {
	//declaration of string and numeric constants
	//good practice for constant variables is to use UPPERCASE letters
	const NAME string = "Juan Dela Cruz"
	const PRICE int = 100

	fmt.Println("Name is", NAME)
	fmt.Println("Price is", PRICE)

	//Declare the integer literals
	//an integer literal can be decimal, octal, or hexadecimal
	// 0x - hex, 0 for octal
	const (
		DECIMAL     = 225  //decimal with no prefix
		OCTAL       = 0377 //octal leading with 0
		HEXADECIMAL = 0xFF //hexadecimal leading with 0x
	)
	fmt.Println("Decimal is", DECIMAL, "Octal is", OCTAL, "Hexadecimal is", HEXADECIMAL)

	//floating number literals
	//can have integer part, decimal point, and exponent
	const PI float64 = 3.141
	const AVOGADRO float64 = 6.022e23 //6.022 * 10^23

	fmt.Println("Pi is", PI)
	fmt.Println("Avogadro's number is", AVOGADRO)

	//escape sequences in string literals
	const GREETING = "Hello guys\n"
	const QUOTE = "	\"Go is simple!\" - A programmer"

	fmt.Println(GREETING)
	fmt.Println(QUOTE)
	//Escape sequences in string literals

	//Alert or Bell
	const BELL = "Bell is \a"
	fmt.Println(BELL)
	//line break
	const LB = "i \nam \ngood"
	fmt.Println(LB)

	//multi-line and concatenated strings
	const MULTILINE = "foobar barfoo" + "farboo boofar"
	const CONCAT = "concat" + "string"
	fmt.Println(MULTILINE)
	fmt.Println(CONCAT)

	//boolean constants
	const ACTIVE = true
	const READY = false

	fmt.Println("Active is", ACTIVE, "Ready is", READY)

	//for calculations
	const LENGTH = 5
	const WIDTH = 10
	const AREA = LENGTH * WIDTH
	fmt.Println("Area is", AREA)
}
