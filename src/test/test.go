package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println("i=", i, s)
		runtime.Gosched()
		fmt.Println(s)

	}
}

func main() {
	fmt.Println("Start ...")

	ch := make(chan string, 10)

	ch <- "jhhh"
	ch <- "done"

	for {
		select {
		case value := <-ch:
			//value := <-ch
			fmt.Println(value)
		}

	}

	go say("world") //开一个新的Goroutines执行
	say("hello")    //当前Goroutines执行
	// ch := make(chan int, 1)
	// countYes := 0
	// countNo := 0
	// for {
	// 	select {
	// 	case ch <- 0:
	// 	case ch <- 1:
	// 	}

	// 	i := <-ch
	// 	if i == 1 {
	// 		countYes++
	// 	} else {
	// 		countNo++
	// 	}

	// 	fmt.Println(countYes, countNo, "Value received:", i)
	// }
}
