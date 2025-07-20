package main

import "fmt"

func main() {
	// integers
	//signed integers are negative and positive integers
	//unsigned integers are only positive integers

	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	//assigning values to integers
	i = -128
	i8 = 127
	i16 = -32767
	i32 = 2147483647
	i64 = -9223372036854775807

	//unassigned integers
	var u uint
	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64

	//assigning values to integers
	u = 255
	u8 = 255
	u16 = 65535
	u32 = 4294967295
	u64 = 18446744073709551615

	fmt.Println("Signed Integers:", i, i8, i16, i32, i64)
	fmt.Println("Unsigned Integers", u, u8, u16, u32, u64)

	//Floating point numbers
	// float32
	var f32 float32 = 10.6

	// and float64 are used for floating point numbers
	var f64 float64 = 10.6

	//printing floating point numbers
	fmt.Println("Floating Point Numbers:", f32)
	fmt.Println("Floating Point Numbers:", f64)

	// demonstrating high precision between float32 and float64
	var HP float64 = 123456789012345
	var LP float32 = 123456789012345

	fmt.Println("High Precision Float64:", HP)
	fmt.Println("Low Precision Float32:", LP)

	// Booleans
	var isActive bool = true
	var isOn bool = false

	fmt.Println("Boolean Values:", isActive)
	fmt.Println("Boolean Values:", isOn)

	//Complex data types

	var CN1 complex128 = complex(5, 10) //first one is real part and second one is imaginary part
	var CN2 complex64 = complex(2, 4)

	//prints without any appending spaces
	//print(CN1)
	//print(CN2)

	fmt.Println("Complex Numbers:", CN1)
	fmt.Println("Complex Numbers:", CN2)

	//Strings
	var name string = "Steve Martin"
	fmt.Println("String Value:", name)
}
