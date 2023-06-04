package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var mut sync.Mutex    // Usually it's a pointer
var wg sync.WaitGroup // Usually it's a pointer

/*
GoRoutines :- Do not communicate by sharing memory; instead; share memory by communicating
*/
func main() {
	// go greeter("Hello")
	// greeter("World")
	websiteList := []string{"https://google.com", "https://fb.com", "https://go.dev", "https://github.com"}
	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

//	func greeter(s string) {
//		for i := 0; i < 6; i++ {
//			time.Sleep(3 * time.Millisecond)
//			fmt.Println(s)
//		}
//	}
func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPs in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
