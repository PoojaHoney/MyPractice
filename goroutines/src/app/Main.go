// creating goroutines
// synchronozation
// 	waitgroups
// 	mutexes
// Parallelism
// Best Practices

// go routines - going to allow our application to run on multiple functionality at a time
//concurrency - ability of the go lang to work on multiple things at the same time

package main

import (
	"fmt"
	// "runtime"
	// "sync"
	"time"
)

// Mutex - is usally the lock that the application is going honor

// var wg = sync.WaitGroup{} //synchronize multiple goroutines at a time
// var m = sync.RWMutex{}
// var counter = 0

// func main() {
// 	runtime.GOMAXPROCS(100) //tells that number of thread available // used on large synchronizing call
// 	fmt.Println(runtime.GOMAXPROCS(-1))
// 	for i := 0; i < 10; i++ {
// 		wg.Add(2) //this is what we are blcoking the go routines to run and making this as single thread by this mutex
// 		m.RLock()
// 		go sayHello()
// 		m.Lock()
// 		go increment()
// 	}
// 	wg.Wait()
// }

// func increment() {
// 	// m.Lock()
// 	counter++
// 	m.Unlock()
// 	wg.Done()
// }

// func sayHello() {
// 	// m.RLock()
// 	fmt.Printf("Hello Babe: %v\n", counter)
// 	m.RUnlock()
// 	wg.Done()
// }

// func main() {
// 	for i := 0; i < 10; i++ {
// 		wg.Add(2) //Here the output is a bit messy with an unexpected result because the go routies fight each other to run over each other to overcome this we can use the Mutex
// 		go sayHello()
// 		go increment()
// 	}
// 	wg.Wait()
// }

func main() {
	// go sayHello() //spin off called a green thread and run the sayHello function in that green thread
	// time.Sleep(100 * time.Millisecond) // time delay to stop the main function routine mean while the say hello go routine can execute

	var msg = "Hello"
	go func() {
		sayHello(msg) //here msg is the copy foe Hello 
	}() // to avoid race conditions
	msg = "Bye Bye" //Racing Condition without parameters to anonymous function
	time.Sleep(100 * time.Millisecond)

	// var msg = "Hello"
	// go func(msg string) {
	// 	sayHello(msg) //here msg is the copy foe Hello 
	// }(msg) // to avoid race conditions
	// msg = "Bye Bye" //Racing Condition without parameters to anonymous function
	// time.Sleep(100 * time.Millisecond)

	//-------------------Weight Groups
	//Alternatives to avoid the time sleep for the production is done by weight groups
	// var msg = "Hello" 
	// wg.Add(1)
	// go func(msg string) {
	// 	sayHello(msg) 
	// 	wg.Done()
	// }(msg) 
	// msg = "Bye Bye" 
	// wg.Wait()
}

// func sayHello() {
// 	fmt.Printf("Hey Babe: %v\n", counter)
// 	wg.Done()
// }

// func increment() {
// 	counter++
// 	wg.Done()
// }

// func sayHello() {
// 	fmt.Println("Hello")
// }

func sayHello(msg string) {
	fmt.Println(msg)
}


//go run -race Main.go - the go complier identifies and locate places in it that have the potential of access seeing the same 
//memory at the same time, or in an unsynchronized way, causing very subtle and disacter bugs in your application