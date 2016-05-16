// 예제 1부터 9까지
/*
RUBY:
Convention over configuration

GO
Configuration, configuration, configuration

RESOURCE SERVER
- listen on a TCP port
- handle requests: route a URL to a file

ServeMux = HTTP request router = multiplexor = mux
http://www.alexedwards.net/blog/a-recap-of-request-handling

Mutexes are something else
http://www.alexedwards.net/blog/understanding-mutexes

NICE TO KNOW
www.rawgit.com
https://cdn.rawgit.com/GoesToEleven/go_web_app/01_working_static_page/templates/home.html

NET/HTTP

- http.ListenAndServe
-- listens for, and responds to, http requests
-- handles each request using go routines
--- lightweight concurrency (eg, coroutines - processes --> threads --> coroutines)
--- this is multiplexing, thus, multiplexor ( = HTTP request router = ServeMux = mux )
--- blacks main thread (call after configuration of server complete)

- http.Handle
-- handles a URL request
-- maps a URL to any TYPE ("object") implementing the handler interface
--- http://golang.org/pkg/net/http/#Handler

- http.HandleFunc
-- handles a URL request
-- maps a URL to a FUNCTION
--- "wrapper" around a function
---- turns any function into a handler

Handle -> handler
HandleFunc -> handlerFunc


*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	//"text/template"
	//"html/template"
)

type person struct {
	fName string
}

func (p *person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("First Name: " + p.fName))
}

type MyHandler struct{}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)
	//    path := "templates" + r.URL.Path

	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 My Friend - " + http.StatusText(404)))
	}
}

type MyHandler2 struct{}

func (this *MyHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)
	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string

		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		} else if strings.HasSuffix(path, ".mp4") {
			contentType = "video/mp4"
		} else {
			contentType = "text/plain"
		}

		w.Header().Add("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 Mi amigo - " + http.StatusText(404)))
	}
}

type MyHandler3 struct{}

func (this *MyHandler3) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)
	f, err := os.Open(path)

	if err == nil {
		bufferedReader := bufio.NewReader(f)

		var contentType string

		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		} else if strings.HasSuffix(path, ".mp4") {
			contentType = "video/mp4"
		} else {
			contentType = "text/plain"
		}

		fmt.Println(contentType)
		w.Header().Add("Content-Type", contentType)
		bufferedReader.WriteTo(w)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 Mi amigo - " + http.StatusText(404)))
	}
}

// http://localhost:8080/templates/home.html

func main() {
	//simpleServer1()
	//simpleServer2()
	//simpleServer3()
	//simpleFileServer1()
	//simpleFileServer2()
	simpleFileServer3()
	//builtinFileServer1()
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello universe"))
}

func simpleServer1() {
	http.HandleFunc("/", someFunc)
	http.ListenAndServe(":8080", nil)
	// nil means use default ServeMux
}

func simpleServer2() {
	myMux := http.NewServeMux()
	myMux.HandleFunc("/", someFunc)
	http.ListenAndServe(":8080", myMux)
}

func simpleServer3() {
	personOne := &person{fName: "Samil Lee"}
	http.ListenAndServe(":8080", personOne)
}

func simpleFileServer1() {
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":8080", nil)
}

func simpleFileServer2() {
	http.Handle("/", new(MyHandler2))
	http.ListenAndServe(":8080", nil)
}

func simpleFileServer3() {
	http.Handle("/", new(MyHandler3))
	http.ListenAndServe(":8080", nil)
}

func builtinFileServer1() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("")))
}

/*
http.HandleFunc("/", someFunction)
Go matches requests to the most specific route registered
http://golang.org/pkg/net/http/#ServeMux
everything matches "/"

nil
- http.ListenAndServe(":8080", nil)
- meaning: use the DefaultServeMux
http://golang.org/pkg/net/http/#pkg-variables

behind the scenes:
- request comes in
- received on primary thread
- goroutine created
-- runs concurrently to main thread
-- lightweight
- request passed to goroutine
- handling multiple requests at same time: "multiplexing"

structuring go apps

*****************
SUPER IMPORTANT
*****************

ROOT is determined by ...
WHERE YOU ARE IN YOUR DIRECTORY WHEN YOU RUN "GO RUN"
wherever you are in your dir when you run "go run" - that's your root

for example
if this is your file structure
go_web_app
-- bin
-- pkg
-- src
---- main
------ main.go
-- templates
---- home.html

if you were in "go_web_app / src / main"
and you ran "go run main.go"
then the root of your files would be "go_web_app / src / main"
and you would not be able to access "templates" nor "templates/home.html"

if you were in "go_web_app"
and you ran "go run main.go"
then the root of your files would be "go_web_app"
and you would be able to access "templates" and "templates/home.html"

this is a good article - though advanced (so don't freak out if you're just getting started)
https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091

this can also be useful in debugging:
-- import "os"
---- log.Println("ENV: ", os.Environ())
*/
