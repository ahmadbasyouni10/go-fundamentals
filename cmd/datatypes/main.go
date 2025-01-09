// main package is what compiler looks for to start the program in the files that are in the main package
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// fmt.Println("Hello World!")
	// var x of type int
	// pointless variables will throw error needs to be used

	// max value of int16 is 32767
	//var x int16 = 32769
	// doing anything above the max value will cause overflow
	var x int16 = 32767
	var y int32 = 10
	fmt.Println(x + 1) // this will cause overflow but wont throw error instead it will give the min value of int16

	var z uint16 = 65535 // max value of uint16 is 65535 bc it can store 0 to 65535 so no negative values (unsigned)
	// int64 takes 2 times the memory of int32 but can store 2 times the value of int32
	fmt.Println(x, y, z)

	var floatNum float32 = 12345678.9
	fmt.Println(floatNum)
	// float32 will round off the value to 7 decimal places so wont store long floats properly

	var floatNum2 float64 = 123456678.9
	fmt.Println(floatNum2)

	// for rgb, we use uint8 bc it can store 0 to 255 and is better than using int which is 64 bits stored in memory
	// color white in rgb
	var rgb = [3]uint8{255, 255, 255}
	fmt.Println(rgb)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	// this will throw error bc we cant add float32 and int32
	// var result float32 = floatNum32 +intNum32

	// what we can do is cast the int32 to float32
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 10
	var intNum2 int = 3
	fmt.Println(intNum1 / intNum2) // this will give 3 bc both are int so it will give int result (rounded down)
	fmt.Println(intNum1 % intNum2)

	// you cant have a string that spans multiple lines
	//var myString string = "Hello
	//World!"

	var myString string = "Hello World!"
	fmt.Println(myString)

	// what you can do is use /n to go to next line
	var myString2 string = "Hello\nWorld!"
	fmt.Println(myString2)

	var myString3 string = "Hello\tWorld!"
	fmt.Println(myString3)

	// you can use backticks to have a string that spans multiple lines
	var myString4 string = "Salah" + " " + "is the best player in the world"
	fmt.Println(myString4)

	// length of string
	fmt.Println("Length of myString4 is", len(myString4))

	// length of special characters might be 2 for instance the gamma symbol
	var myString5 string = "γ"
	fmt.Println("Length of myString5 is", len(myString5))

	// to fix that we use utf8 package
	fmt.Println("Length of myString5 is", utf8.RuneCountInString(myString5))

	var myRune rune = 'a'
	// rune is for characters and is an alias for int32
	// this will give the ascii value of a
	fmt.Println(myRune)

	var myBoolean bool = true
	fmt.Println(myBoolean)

	// default value of int is 0 and for all ints floats runes and unsigned is 0
	var intNum3 int
	fmt.Println(intNum3)

	// default value of string is empty string
	var myString6 string
	fmt.Println(myString6)

	// default value of bool is false
	var myBoolean2 bool
	fmt.Println(myBoolean2)

	//we can omit the type and let go infer the type
	var myString7 = "Hello World!"
	fmt.Println(myString7)

	// we can also do this
	myString8 := "Hello World!"
	fmt.Println(myString8)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	// using type when declaring vars is good bc if calling a function we wont know unless going to the function
	// and checking the type of the vars
	// so its better to use var1 int = foo() instead of var1 := foo()

	//const is used to declare constants
	const myConst int = 10
	fmt.Println(myConst)

	// good for initializing pi for instance
	const pi float64 = 3.14159
	fmt.Println(pi)

}
