package main

// Parallelize a web crawler without fetching the same URL twice.
// Hint: you can keep a cache of the URLs that have been fetched on a map
// Maps alone are not safe for concurrent use!

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func Crawl2(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	results2.Lock()
	if _, ok := results2.m[url]; ok {
		fmt.Println("\t", url, "checked already")
		results2.Unlock()
		return
	}

	results2.m[url] = &fakeResult{"", nil} // added new
	results2.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	results2.Lock()
	results2.m[url] = &fakeResult{body, urls} // added new
	results2.Unlock()

	done := make(chan bool)

	for _, u := range urls {
		go func(url string) {
			Crawl2(url, depth-1, fetcher)
			done <- true
		}(u)
	}

	for _, _ = range urls {
		<-done
	}
}

func main() {
	//results = make(fakeFetcher)
	//Crawl("http://golang.org/", 4, fetcher)
	Crawl2("http://golang.org/", 4, fetcher)

	//go Crawl2("http://golang.org/", 4, fetcher)

	/*
		done := make(chan bool)
		go func() {
			Crawl2("http://golang.org/", 4, fetcher)
			done <- true
		}()
		<-done
	*/
}

type fakeResult struct {
	body string
	urls []string
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	time.Sleep(time.Second)
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// Crawl* 메쏘드들이 interface 라서 Fetch로만 받아야 됨
func (f fakeFetcher) FetchMe(url string) (string, []string, error) {
	time.Sleep(time.Second)
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

type visits struct {
	sync.Mutex
	m map[string]*fakeResult
}

// if body is empty then searched but not found.
var results2 visits = visits{m: make(map[string]*fakeResult)}
var results fakeFetcher = make(map[string]*fakeResult)

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
