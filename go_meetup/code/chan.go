package main

import (
	"fmt"
	"time"
)

func strGenerator(input string, output chan string, wait time.Duration) {
	for {
		output <- input
		time.Sleep(wait)
	}
}

func main() {
	chA := make(chan string)
	chB := make(chan string)
	done := make(chan bool)
	quit := false
	var str string
	go strGenerator("a", chA, 1000 * time.Millisecond)
	go strGenerator("A", chA, 800 * time.Millisecond)
	go strGenerator("B", chB, 600 * time.Millisecond)
	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()
	for quit == false{
		select {
		case str = <-chA:
			fmt.Println("Message from chan A:", str)
		case str = <-chB:
			fmt.Println("Message from chan B:", str)
		case <-done:
			fmt.Println("Done.")
			quit = true
		}

	}
}
