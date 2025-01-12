// pointers -> variables which store memeory address of another variable

package main

import "fmt"

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("The memory address is: %p\n", &thing2)
	fmt.Printf("the pointer points at: %p\n", thing2)
	for i := range *thing2 {
		(*thing2)[i] = (*thing2)[i] * (*thing2)[i]
	}
	return *thing2
}

func main() {
	//var p *int32              // default value is nil, pointer to an int32
	var i int32 = 42 // default value is 0
	//var g *int32 = new(int32) // creates a pointer to an int32 with a value of 0, finds a empty memory address and assigns it to g

	//fmt.Println("Value of i: ", i)
	//fmt.Println("Value of p: ", p) // this is the memory address of i but since i is not assigned to p, it is nil

	//fmt.Println(&p) // memory address of p
	//fmt.Println(*p) // will throw an error bc p is nil and has no value

	//fmt.Printf("The value that p points to: %v\n", *g) // will print 0 bc the value of g is 0 since it was created with new(int32)

	// also trying to assign the *p to a value will throw an error bc p is nil
	//*p = 8

	var z *int32 = new(int32)
	z = &i
	fmt.Println(*z)
	*z = 8
	// the value of i is changed to 8 since z is assigned to the memory address of i
	fmt.Println(i)

	// this doesnt apply when using normal variables
	var x int32 = 43
	var y int32 = x
	y = 9
	fmt.Println(x, y)

	// slices have a pointer to the array they are pointing to so changing the slice will change the array as well
	var slice = []int32{1, 2, 3}
	var slice2 = slice
	slice2[0] = 9
	fmt.Println(slice, slice2)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory address is: %p\n", &thing1)

	var result [5]float64 = square(&thing1)
	fmt.Println(result)
	// passing an array to a function will create a copy of the array and not the pointer to the array so changing the array in the function will not change the original array

	// way more memory used when passing arrays to functions so instead use pointers

}

// the &thing2 is the memory address of the pointer to the original array so wont be the same as the memory address of the original array
// the thing2 is the memory address of the original array and thats the value which is the pointer to the original array
// the thing1 is the original array and the *thing2 is the original array because * is the dereference operator which gives us the value of the pointer
