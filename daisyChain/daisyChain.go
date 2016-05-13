package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 1000000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}
