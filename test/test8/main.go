package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) // channel

	go hello("hello", ch) // goroutine, async work
	for i := range ch {   //listening channel, not work without channel closing, drop error
		fmt.Println(i)
	}

	ch2 := make(chan string)
	go sayBark(ch2)
	<-ch2

	data := make(chan int)
	exit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-data)
		}
		exit <- 0
	}() // declaring and calling an anonymous function
	selectOne(data, exit)
}

func hello(str string, ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(10 * time.Millisecond)
		ch <- i
	}
	close(ch) // close channel
}

func say(word string) {
	time.Sleep(1 * time.Second)
	fmt.Println(word)
}

func sayBark(exit chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(10 * time.Millisecond)
		say("Bark")
	}
	exit <- "s"
}

func selectOne(data, exit chan int) { // work with multiple channels
	x := 0
	for {
		select {
		case data <- x:
			x += 1
		case <-exit:
			fmt.Println("exit")
			return
		}
	}
}
