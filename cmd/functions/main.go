package main

import (
	"errors"
	"fmt"
)

func main() {
	var value string = "Hello World!"
	printMe(value)
	var n1 int = 10
	var n2 int = 5
	var result, result2, err = intDivision(n1, n2)
	// %d is for integers placeholder in the string and %s is for string placeholder, %v is for any type of value
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	} else if result2 == 0 {
		fmt.Printf("Result: %d\n", result)
	} else {
		fmt.Printf("Result: %d, Rema inder: %d\n", result, result2)
	}
	/*
		switch {
		case err != nil:
			fmt.Printf("Error: %s\n", err)
			// break is implied
		case result2 == 0:
			fmt.Printf("Result: %d\n", result)
		default:
			fmt.Printf("Result: %d, Remainder: %d\n", result, result2)
		}
	*/

	switch result2 {
	case 0:
		fmt.Printf("The division was exact")
	case 1, 2:
		fmt.Printf("The division was very close to being exact")
	default:
		fmt.Printf("The division was not close")
	}

}

func printMe(value string) {
	fmt.Println(value)
}

// two return types (int, int)
func intDivision(num int, den int) (int, int, error) {
	// default value of nil
	var err error
	if den == 0 {
		// panic will stop the program and print the message
		// panic("Cannot divide by zero")
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = num / den
	var result2 int = num % den
	// returns the result and the remainder
	return result, result2, err
}

 