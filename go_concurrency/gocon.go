package main

import (
	"fmt"
	"math/rand"
	"time"
	//"image"
	//"io"
	//"math"
	//"runtime"
	//"strings"
	//"sync"
)

const Pi = 3.14
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func main() {
	const World = "삼일"

	go boring("without Channnel")

	bChan := make(chan string)
	go boring2("with Channel", bChan)
	for i := 0; i < 5; i++ {
		fmt.Printf("%v\n", <-bChan)
	}

	bChan2 := boring3Pattern("Generation pattern: Channel return")
	for i := 0; i < 5; i++ {
		fmt.Printf("%v\n", <-bChan2)
	}

	bChan3 := fanIn(boring3Pattern("multiplexing random no select"), boring3Pattern("Go 3333"))
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", <-bChan3)
	}

	bChan4 := fanInSelect(boring3Pattern("multiplexing random with select"), boring3Pattern("Go 4444"))
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", <-bChan4)
	}

	fmt.Println("")
	bChan33 := fanInSeq(boringSeq("multiplexing seq 1"), boringSeq("multiplexing seq 2"), boringSeq("multiplexing seq 3")) // HL
	for i := 0; i < 5; i++ {
		msg1 := <-bChan33
		fmt.Println(msg1.str)

		msg2 := <-bChan33
		fmt.Println(msg2.str)

		msg3 := <-bChan33
		fmt.Println(msg3.str)

		msg1.wait <- false // false 라도 작동, 단지 하나씩 처리 되도록 하는 것에 의미가 더 큼
		msg2.wait <- true
		msg3.wait <- true
	}

	fmt.Println("")
	runWithTimeout("select with timeout for each run", 10, true)

	fmt.Println("")
	runWithTimeout("select with global timeout", 10, false)

	fmt.Println("")
	quit := make(chan string)
	bChan7 := boring4Pattern("Select with quit channel", quit)
	//for i := rand.Intn(10); i >= 0; i-- {
	for i := 10; i >= 0; i-- {
		fmt.Println(<-bChan7)
	}
	quit <- "byte"
	fmt.Println("Quit:", <-quit)

}

// Go Concurrency Patterns  Google I/O 2012
func boring(msg string) {
	le3 := 10
	for i := 0; i < 5; i++ {
		n := time.Now()
		time.Sleep(time.Duration(rand.Intn(le3)) * time.Millisecond)
		fmt.Println(msg, i, time.Since(n))
	}
}

func boring2(msg string, c chan string) {
	//r := rand.NewSource(time.Now().UnixNano())
	le3 := 10
	for i := 0; ; i++ {
		n := time.Now()
		time.Sleep(time.Duration(rand.Intn(le3)) * time.Millisecond)
		c <- fmt.Sprintf("%s %d %v", msg, i, time.Since(n))
	}
}

func boring3Pattern(msg string) <-chan string { // returns receive-only channel
	c := make(chan string)

	go func() { // launch the goroutine
		le3 := 10
		for i := 0; ; i++ {
			n := time.Now()
			time.Sleep(time.Duration(rand.Intn(le3)) * time.Millisecond)
			c <- fmt.Sprintf("%s %d %v", msg, i, time.Since(n))
		}
	}()
	return c
}

// rcvquit
func boring4Pattern(msg string, quit chan string) <-chan string { // returns receive-only channel
	c := make(chan string)

	go func() { // launch the goroutine
		le3 := 10
		for i := 0; ; i++ {

			n := time.Now()
			time.Sleep(time.Duration(rand.Intn(le3)) * time.Millisecond)

			select {
			case c <- fmt.Sprintf("%s %d %v", msg, i, time.Since(n)):
				// do nothing
			case <-quit:
				cleanup()
				quit <- "See you"
				return
			}
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func fanInSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func cleanup() {

}

// Sequencing
type MessageSeq struct {
	str  string
	wait chan bool
}

func boringSeq(msg string) <-chan MessageSeq {
	c := make(chan MessageSeq)

	waitForIt := make(chan bool) // Shared between all messages.

	go func() {
		for i := 0; ; i++ {
			c <- MessageSeq{fmt.Sprintf("%s: %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c // Return the channel to the caller. // HL
}

func fanInSeq(inputs ...<-chan MessageSeq) <-chan MessageSeq { // HL
	c := make(chan MessageSeq)
	for i := range inputs {
		input := inputs[i] // New instance of 'input' for each loop.
		go func() {
			for {
				c <- <-input
			}
		}()
	}
	return c
}

func runWithTimeout(msg string, ms uint, forEach bool) {

	c := boring3Pattern(msg)
	//bChan6 := boring3Pattern("Select global timeout")

	if forEach {
		for i := 0; i < 5; i++ {
			//for { // 이럼 영원히 안빠져나올 가능성 있음
			select {
			case s1 := <-c:
				fmt.Println(s1)
			case <-time.After(time.Duration(ms) * time.Millisecond): // 각각이 timeout 해당 되는지 점검
				fmt.Println("--- timeout", time.Duration(ms)*time.Millisecond)
				return
			}
		}
	} else {
		timeout := time.After(time.Duration(ms) * time.Millisecond)
		for {
			select {
			case s1 := <-c:
				fmt.Println(s1)
			case <-timeout: // 전체에 대해 위 지정 시간
				fmt.Println("--- timeout (return)", time.Duration(ms)*time.Millisecond)
				return
			}
		}
	}
}
