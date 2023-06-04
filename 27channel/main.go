package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang -- LearnCodeOnline")

	mych := make(chan int, 2)
	wg := &sync.WaitGroup{}
	// mychan <- 5
	// fmt.Println(<-mychan)
	wg.Add(2)

	// Read Only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch // To identify "0" is a signal or for close channel
		//fmt.Println(<-ch)
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		//fmt.Println(<-ch)
		wg.Done()
	}(mych, wg)
	// Write Only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 0 // Its a signal
		// ch <- 5
		// ch <- 6
		close(ch) // Gives 0 at the output
		wg.Done()
	}(mych, wg)
	wg.Wait()
}
