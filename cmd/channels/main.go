package main

import (
	"fmt"
	"time"
)

func main() {
	// unbuffered channel
	//var c = make(chan int)
	// buffered channel with space for 5 elements
	// this is helpful when we have a consumer and producer and the producer is faster than the consumer
	// so the producer can keep adding values to the channel and the consumer can keep taking values from the channel on its own pace
	// if using unbuffered channel the producer needs to wait for the main function to read the value from the channel so that it can be popped from the channel and the producer can add another value
	var d = make(chan int, 5)
	//c <- 1      // think of channel as a pipe (array) unbuffered channel space for 1 element
	//var i = <-c // pops the element from the channel and stores it in i, if we print i we get deadlock error write to unbuffered channel code blocked ubtil something reads from it
	// so basically get stuck at c <- 1 and never reaches var i = <-c
	//fmt.Println(i)

	// to use properly we need to use go routines
	// this is a buffered channel with space for 1 element

	// with this now the code that would block code and cause deadlock is in the process function and go routines are used to run the process function in a separate thread
	// so the main function can continue to run and when we get to the fmt.Println(<-c) the code will block until the process function sends a value to the channel (value has been set)
	go process(d)
	//fmt.Println(<-c)
	// we need a for loop to get all the values from the channel

	// keeps going until the channel is closed
	for i := range d {
		fmt.Println(i)
		time.Sleep(time.Second)
	}

}

func process(c chan int) {
	// this runs right before the function returns
	defer close(c)
	// c <- 123
	for i := 0; i < 5; i++ {
		c <- i
	}
	// if we dont close the channel we get a deadlock error cuz main func still listening for values
	// close(c)
	// we can do it with defer statement

	// producer can finish and close the channel and the consumer can keep reading from the channel until the channel is closed which is when using buffered channels
	fmt.Print("Process done")
}
