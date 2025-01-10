package main

import "fmt"

func main() {
	var intArr [3]int32
	fmt.Println(intArr[0])
	fmt.Println(intArr[1])
	fmt.Println(intArr[2])

	// int32 is 4 bytes and array stores 3 int32 so 3*4 = 12 bytes
	// the & gives us the address of the variable in memory and they are contiguous
	// benefit: copmpiler doesnt need to store the address of each element in the array
	// just know the first and increment by 4 bytes to get the next element

	// the address of the first element in the array
	fmt.Println(&intArr[0])
	// the address of the second element in the array
	fmt.Println(&intArr[1])
	// the address of the third element in the array
	fmt.Println(&intArr[2])

	// non exclusive range
	intArr[1] = 10
	fmt.Println(intArr[1:3])

	// to assign an array in initialization
	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	// easier way to assign an array in initialization
	intArr3 := [3]int32{1, 2, 3}
	// or
	intArr4 := [...]int32{1, 2, 3}

	fmt.Println(intArr3)
	fmt.Println(intArr2)
	fmt.Println(intArr4)

	// slices are wrappers around arrays, they are arrays with extra functionality
	var intSlice []int32 = []int32{1, 2, 3}
	fmt.Println(intSlice)
	fmt.Println(len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Println(len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlicenew := append(intSlice, intSlice2...)
	fmt.Println(intSlicenew)

	// pre allocating can slow down the program bc it needs to copy the elements to the new array when the capacity is reached
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	// slides have a length and a capacity
	// length is the number of elements in the slice
	// capacity is the number of elements in the underlying array starting from the first element of the slice
	// the capacity is always a power of 2 and is always greater than or equal to the length
	// originally the capacity is the same as the length but when you append an element, the capacity doubles if the length is equal to the capacity already

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
	myMap["one"] = 1
	fmt.Println(myMap)

	// to delete a key value pair
	delete(myMap, "one")
	fmt.Println(myMap)

	// easier way to assign a map in initialization
	var myMap2 = map[string]uint8{"Adam": 1, "Eve": 2}
	fmt.Println(myMap2)

	// to check if a key exists
	value, exists := myMap2["Adam"]
	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("Key does not exist")
	}
}
