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
		//fmt.Printf("%q\n", <-bChan)
		fmt.Printf("%v\n", <-bChan)
	}

	bChan2 := boring3Pattern("Generation pattern: Channel return")
	for i := 0; i < 5; i++ {
		fmt.Printf("%v\n", <-bChan2)
	}

	bChan3 := fanIn(boring3Pattern("multiplexing"), boring3Pattern("Go 3333"))
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", <-bChan3)
	}

	bChan4 := fanIn2(boring3Pattern("multiplexing with select"), boring3Pattern("Go 4444"))
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", <-bChan4)
	}

	bChan5 := boring3Pattern("select timeout for each run")
	for i := 0; i < 3; i++ {
		//for { // 이럼 영원히 안빠져나옴 각각의 시간초과가 1초를 넘지 않기 때문
		select {
		case s1 := <-bChan5:
			fmt.Println(s1)
		case <-time.After(1 * time.Second): // 각각이 최대 1초
			fmt.Println("--- timeout")
			return
		}
	}

	bChan6 := boring3Pattern("Select global timeout")
	timeout := time.After(50 * time.Millisecond)
	for {
		select {
		case s1 := <-bChan6:
			fmt.Println(s1)
		case <-timeout: // 저체가 최대 위 지정 시간
			fmt.Println("--- timeout")
			return
		}
	}

	quit := make(chan string)
	bChan7 := boring4Pattern("Select with quit channel", quit)
	//for i := rand.Intn(10); i >= 0; i-- {
	for i := 5; i >= 0; i-- {
		fmt.Println(<-bChan7)
	}
	quit <- "byte"
	fmt.Println("Quit:", <-quit)

}

// Go Concurrency Patterns  Google I/O 2012
func boring(msg string) {
	//r := rand.NewSource(time.Now().UnixNano())
	le3 := 10
	for i := 0; i < 5; i++ {
		n := time.Now()
		//time.Sleep(time.Duration(rand.Intn(le3)) * time.Millisecond)
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

func fanIn2(input1, input2 <-chan string) <-chan string {
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

func cleanup() {

}

// Sequencing
type Message struct {
	str  string
	wait chan bool
}
