// channels basics
// restrict data flow
// buffered channels
// closing channels
// for range loops
// select statements

// Channels are mainly designed for the synchornize data transmission between multiple go routines

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 2)

	wg.Add(2)
	//receving go routine
	go func(ch <- chan int) {
		i := <-ch
		fmt.Println(i)
		// ch <- 27 // this shows the error because this is the receving channel only
		wg.Done()
	}(ch)
	//sending go routine
	go func(ch chan <- int) {
		ch <- 42 //copy of the data
		// fmt.Println(<-ch)// this is an error because this only the send only channel
		wg.Done()
	}(ch)
	wg.Wait()

	// wg.Add(2)
	// //receving go routine
	// go func() {
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	ch <- 27
	// 	wg.Done()
	// }()
	// //sending go routine
	// go func() {
	// 	ch <- 42 //copy of the data
	// 	fmt.Println(<-ch)
	// 	wg.Done()
	// }()
	// wg.Wait()

	// for i := 0; i < 5; i++ {
	// 	wg.Add(2)
	// 	//receving go routine
	// 	go func() {
	// 		i := <-ch
	// 		fmt.Println(i)
	// 		wg.Done()
	// 	}()
	// 	//sending go routine
	// 	go func() {
	// 		ch <- 42 //copy of the data
	// 		wg.Done()
	// 	}()
	// 	wg.Wait()
	// }

	//---------------------deadlock
	// receving go routine
	// go func() {
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	wg.Done()
	// }()
	// for i := 0; i < 5; i++ {
	// 	wg.Add(2)
	// 	//sending go routine
	// 	go func() {
	// 		ch <- 42 //copy of the data
	// 		wg.Done()
	// 	}()
	// 	wg.Wait()
	// }

	// wg.Add(2)
	// // receving go routine
	// go func(ch <-chan int) {
	// 	// i := <-ch
	// 	// fmt.Println(i)
	// 	// i = <-ch
	// 	// fmt.Println(i)
	// 	// wg.Done()
	// 	for i := range ch {
	// 		fmt.Println(i)
	// 	}
	// 	wg.Done()
	// }(ch)
	// //sending go routine
	// go func(ch chan<- int) {
	// 	ch <- 42 //copy of the data
	// 	ch <- 27
	// 	close(ch)
	// 	wg.Done()
	// }(ch)
	// wg.Wait()

	// wg.Add(2)
	// go func(ch <-chan int) {
	// 	for {
	// 		if i, ok := <-ch; ok {
	// 			fmt.Println(i)
	// 		} else {
	// 			break
	// 		}
	// 	}
	// 	wg.Done()
	// }(ch)

	// go func(ch chan<- int) {
	// 	ch <- 42
	// 	ch <- 23
	// 	fmt.Println(ch)
	// 	close(ch)
	// 	wg.Done()
	// }(ch)
}
