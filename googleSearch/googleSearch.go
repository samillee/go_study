package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	//results := Google("golang")
	//results := Google2("golang")
	results := Google21("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)

}

type Result struct {
	url string
}

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

func Google2(query string) (results []Result) {
	c := make(chan Result)

	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		r := <-c
		results = append(results, r)

	}
	return
}

func Google21(query string) (results []Result) {
	c := make(chan Result)

	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		//case r := c:
		case r := <-c:
			results = append(results, r)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

/*
func Google3(query string) (results []Result) {
	c := make(chan Result)

	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2) }()
	go func() { c <- First(query, Video1, Video2) }()

	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case r := <-c:
			results = append(results, r)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}
*/
