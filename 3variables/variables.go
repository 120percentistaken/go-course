package main

import "fmt"

func main() {
	//1st way: Declare and assign on 2 different lines

	var apple string = "This is a big apple"
	var weight int = 65 //in kilograms

	//2nd way: in 1 line
	var height int = 24
	fmt.Println("apple:", apple)
	fmt.Println("weight:", weight)
	fmt.Println("height:", height)

	//3rd way: short handed declaration
	//Type inference is done by the compiler
	// to assign types to entities that otherwise lack any type annotations.
	// The compiler effectively just 'fills in' the static type information
	// on behalf of the programmer.
	age := 26 //in years
	country := "Philippines"
	fmt.Println("My age is ", age)
	fmt.Println("I live in the", country)

	var chocolate, honey int = 10, 20 //multiple variables in one line
	fmt.Println("I have", chocolate, "chocolates and", honey, "honey jars")

	//combining variables
	//to get the total number of sweets
	var sweets = chocolate + honey
	fmt.Println("Total number of sweets is", sweets)

	var windows, mac, linux string = "Windows is ok\n", "Mac is great\n", "Linux is awesome\n"
	fmt.Println(windows, mac, linux)

	//static type declaration
	var x float64 = 20.5
	fmt.Println(x)
	//another function imported from format package
	fmt.Printf("x is is of type: %T\n", x)

	//dynamic type declaration
	y := 20.5
	fmt.Println(y)
	fmt.Printf("y is is of type: %T\n", y)
	//mixed variable declaration
	var a, b, c = 23, 3, "foobar"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("a is is of type: %T\n", a)
	fmt.Printf("b is is of type: %T\n", b)
	fmt.Printf("c is is of type: %T\n", c)
}
