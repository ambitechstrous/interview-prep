package main

import (
	"fmt"
	"sync"
	"test-golang-concurrency/internal/atomic"
	"test-golang-concurrency/internal/client"

	"github.com/google/uuid"
)

func doSomeStuff(wg *sync.WaitGroup, num *int) int {
	defer wg.Done()
	*num = *num + 1
	return *num
}

// The below produces a race condition on purpose
// func main() {
// 	fmt.Println("Starting concurrency test...")
// 	num := 100

// 	var wg sync.WaitGroup
// 	for i := 0; i < 100; i++ {
// 		wg.Add(1)
// 		go func() {
// 			res := doSomeStuff(&wg, &num)
// 			fmt.Printf("found number %d\n", res)
// 			fmt.Printf("found number %d\n", num)
// 		}()
// 	}

// 	// Wait for everything to pass
// 	wg.Wait()
// 	fmt.Println(num)
// }

func main() {
	fmt.Println("Starting concurrency test...")
	num := atomic.NewAtomicInteger(100)

	// Start with working group example
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			num.Add(1)
			fmt.Printf("found number %d\n", num.Get())
		}()
	}

	// Wait for everything to pass
	wg.Wait()
	fmt.Println(num.Get())

	// Set up mocked database client for concurrency
	database := client.NewConcurrentClient()

	// Test unbuffered channel (just a single done command)
	done := make(chan bool)

	go func() {
		database.WriteData("hello", "world")
	}()

	go func() {
		database.WriteData("testing", "hello")
	}()

	go func() {
		database.WriteData("hello", "other")
		done <- true
	}()

	fmt.Printf("hello: %s\n", database.ReadData("hello"))
	<-done

	fmt.Printf("hello: %s\n", database.ReadData("hello"))
	fmt.Printf("testing: %s\n", database.ReadData("testing"))

	// Test buffered channel (wait for multiple commands)
	bufferSize := 5

	// Write a bunch of data. Buffer keys via channel to retrieve from go routine.
	keyChan := make(chan string, bufferSize)
	for i := 0; i < bufferSize; i++ {
		go func() {
			key := uuid.New().String()
			data := fmt.Sprintf("value for key %s", key)
			database.WriteData(key, data)
			keyChan <- key
		}()
	}

	// Channel to buffer in values from keys
	buffered := make(chan string, bufferSize)
	for i := 0; i < bufferSize; i++ {
		go func() {
			key := <-keyChan
			data := database.ReadData(key)
			buffered <- data
		}()
	}

	// Retrieve data from buffer
	for i := 0; i < bufferSize; i++ {
		data := <-buffered
		fmt.Printf("got data %s\n", data)
	}
}
