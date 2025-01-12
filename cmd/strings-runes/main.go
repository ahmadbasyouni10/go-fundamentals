package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "rèsumè"
	fmt.Println(myString)
	// returns 8 bc the string is stored in utf-8 and the è is stored in 2 bytes
	fmt.Println(len(myString))

	// wont work bc the string is stored in utf-8 and indexing will give us the byte value if it was stored in ascii or utf-32
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)

	// will work bc range gives us proper values for utf-8
	// thats why index 2 is skipped bc the è is stored in 2 bytes
	for i, v := range myString {
		fmt.Println(i, v)
	}

	// this is why using runes is better bc alias for int32
	// now index 2 is not skipped
	var myStringg = []rune("rèsumè")
	for i, v := range myStringg {
		fmt.Println(i, v)
	}
	// will give 232 bc è is stored in 2 bytes properly
	fmt.Println(myStringg[1])

	var myRune = 'a'
	fmt.Println(myRune)

	var res = ""
	var strSlice = []string{"r", "e", "s", "u", "m", "è"}

	for _, v := range strSlice {
		// creating a new string every time is inefficient since cant mutate strings
		res += v
	}

	fmt.Println(res)

	// strings are immutable so this is inefficient
	// res[0] = 'R'

	// a fix for the above problem
	// use

	var res2 strings.Builder
	for _, v := range strSlice {
		// what happens here is a array is allocated and values are appended, then at the end the array is converted to a string much faster
		res2.WriteString(v)
	}
	fmt.Println(res2)
	fmt.Println(res2.String())
}
