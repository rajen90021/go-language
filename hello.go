package main

import (
	"fmt"
	"time"
)

// Function that runs concurrently as a goroutine
func count(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// Create two goroutines
	go count("goroutine 1")
	go count("goroutine 2")

	// Allow goroutines to execute
	time.Sleep(time.Second * 3)
}
