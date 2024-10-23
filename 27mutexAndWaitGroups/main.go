package main

import (
	"fmt"
	"sync"
	"time"
)

//var counter int
//
//func increment() {
//	counter++
//}
//
//func main() {
//	fmt.Println("Mutex and WaitGroup")
//	for i := 0; i < 1000; i++ {
//		go increment() // Multiple Goroutines modifying the same shared variable
//	}
//	time.Sleep(1 * time.Second)
//	fmt.Println("Final counter value:", counter) // Output is unpredictable
//
//}

var counter int
var wg sync.WaitGroup
var mutex sync.RWMutex

func increment() {
	defer wg.Done()
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func main() {
	fmt.Println("Mutex and WaitGroup")
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment() // Multiple Goroutines modifying the same shared variable
	}
	wg.Wait()
	time.Sleep(1 * time.Second)
	mutex.RLock()                                // Allow time for Goroutines to finish
	fmt.Println("Final counter value:", counter) // Output is unpredictable
	mutex.RUnlock()

}
