package main

import (
	"fmt"
	"sync"
	"time"
)

// concurency is the ability to run multiple tasks at the same time but not necessarily at the same time which is called parallelism we get some sorto of parallelism with concurrency
// using go routines we can run multiple tasks at the same time if we have multiple cores (usually 4-8 cores)
// go routines are lightweight threads that are managed by the go runtime

// slice of strings since able to grow and shrink where arrays [5]string cant grow or shrink and are fixed in size, slices have a pointer to the array they are pointing to so changing the slice will change the array as well
var dbData = []string{"data1", "data2", "data3", "data4", "data5"}
var results = []string{}
var mutex = sync.RWMutex{}

// stands for mutual exclusion, used to lock a memory location so only one go routine can access it at a time
var wg = sync.WaitGroup{}

func main() {
	// time Now() returns the current time
	t0 := time.Now()
	// this will run the dbCall function for each element in the dbData slice
	for i := range dbData {
		// when we spawn a go routine we increment the wait group counter
		wg.Add(1)
		// putting go by itself wont work since the main function will finish before the go routines finish concurrently wont wait for the db call and so fast it will make main finish before the go routines finish
		go dbCall(i)
	}
	// wait() will wait for all the go routines to finish before continuing (wait for all the go routines to finish before the main function finishes (counters are all 0))
	wg.Wait()
	// time Since() returns the time passed since the time passed to it
	fmt.Printf("Time taken: %v\n", time.Since(t0))
	fmt.Println("The results are: ", results)
}

// this function will simulate a db call by sleeping for a random amount of time and then printing the result from the db
func dbCall(i int) {
	// rand.Float32() returns a random float32 between 0 and 1 so we multiply it by 2000 to get a random float32 between 0 and 2000
	var delay = 2000
	// time.Duration() converts the float32 to a time.Duration type which is a type that represents a duration of time in nanoseconds
	var val = time.Duration(delay)
	// time.Millisecond is a time.Duration type so multiplying it by val will convert val to milliseconds
	// time.Sleep() will sleep for the amount of time passed to it simulate latency in the db call
	// if we put the lock here it will defeat the purpose of the go routines since it will make the go routines run one after the other
	// the delay will be waited out for each go routine and 6s wuill be the time taken for all the go routines to finish no concurrency
	time.Sleep(val * time.Millisecond)
	// print the result from the db
	//fmt.Println("The result from the db is: ", dbData[i])
	// this will be altered by the go routines multiple threads modifying same memory location (two processes trying to write to the same memory location same time causes memory corruption)
	//mutex.Lock() // so when a routine reaches this lock, a check is made to make sure that a lock was made by another sub routine it will wait for the unlcok and lock the memory itself for the res slice and unlock so other routines can use it as needed'
	// this will make it so when a thread is writing to the memory location no other thread can access it go handles this for us
	//results = append(results, dbData[i])
	//mutex.Unlock()
	// when the go routine finishes we decrement the wait group counter (decrementing the counter will signal the wait group that the go routine has finished)

	// this will call a func that has a lock so only one go routine can access the memory location at a time
	save(dbData[i])
	// this will call a func that has a read lock so multiple go routines can access the memory location at the same time
	log()
	wg.Done()
}

// We can make these calls run concurrently by using go routines (go keyword followed by a function call)
// to avoid the main function from finishing before the go routines finish we use sync package to create a wait group
// wait groups are counters that are incremented when a go routine is started and decremented when a go routine finishes

// with go routines went from 6 seconds to 2 seconds and we had delays to simulate db calls so shows how much faster go routines are than normal functions
// go routines are lightweight threads that are managed by the go runtime so we dont have to manage them ourselves like in other languages like C or C++ where we have to manage threads ourselves
// we need to be careful when using go routines since they can cause memory corruption if we are not careful
// we can use mutexes to lock the memory location so only one go routine can access it at a time
// go provides a read write mutex which allows multiple go routines to read from the memory location but only one go routine can write to it at a time, check other file for that\

func save(data string) {
	mutex.Lock()
	results = append(results, data)
	mutex.Unlock()
}

func log() {
	mutex.RLock()
	fmt.Printf("The results are %v\n", results)
	mutex.RUnlock()
}

// - Performance isn't limited by the number of CPU cores
// - Many goroutines can make progress simultaneously even with fewer cores
// - This is because goroutines spend most time waiting for I/O, not using CPU
// - Adding more goroutines beyond core count can improve throughput

// For CPU-bound tasks (like complex calculations, data processing):
// - Performance is directly limited by available CPU cores
// - Only N goroutines can truly execute in parallel (N = number of cores)
// - Additional goroutines beyond core count must wait their turn
// - Performance improvement scales roughly linearly with cores up to CPU limit
