package main

import (
	"fmt"

	"github.com/samillee/go_study/stringutil"
)

func main() {
	fmt.Printf("hello, world\n")
	a := "이삼일"
	fmt.Printf(stringutil.Reverse(a))
}
