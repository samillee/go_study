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

NET/HTTP
- http.ListenAndServe
-- listens for, and responds to, http requests
-- handles each request using go routines
- http.Handle
-- handles a URL request
-- maps a URL to any TYPE ("object") implementing the handler interface
--- http://golang.org/pkg/net/http/#Handler
--- type Handler interface { ServeHTTP(ResponseWriter, *Request) }
- http.HandleFunc
-- handles a URL request, maps a URL to a FUNCTION

Note the the differencies
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

// http.Handle("/", new(MyHandler))
// 메쏘드 이름은 반드시 ServeHTTP 이어야 함, http.Handle가 interface를 받기 때문
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
	simpleServer1()
	simpleServer2()
	simpleServer3()
	simpleFileServer1()
	simpleFileServer2()
	simpleFileServer3()
	builtinFileServer1()
}

// HandleFunc 함수에 사용되려면 시그너쳐가 중요, 주의 Handle은 전혀 다른 것임
// http.HandleFunc("/", someFunc)
func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello universe"))
}

func simpleServer1() {
	http.HandleFunc("/", someFunc)
	http.ListenAndServe(":6001", nil)
	// nil means use default ServeMux
}

func simpleServer2() {
	myMux := http.NewServeMux()
	myMux.HandleFunc("/", someFunc)
	http.ListenAndServe(":6002", myMux)
}

func simpleServer3() {
	personOne := &person{fName: "Samil Lee"}
	http.ListenAndServe(":6003", personOne)
}

func simpleFileServer1() {
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":6004", nil)
}

func simpleFileServer2() {
	http.Handle("/", new(MyHandler2))
	http.ListenAndServe(":6005", nil)
}

func simpleFileServer3() {
	http.Handle("/", new(MyHandler3))
	http.ListenAndServe(":6006", nil)
}

func builtinFileServer1() {
	http.ListenAndServe(":6007", http.FileServer(http.Dir("")))
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
- goroutine created and runs concurrently to main thread
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

if you were in "go_web_app / src / main" and you ran "go run main.go"
then the root of your files would be "go_web_app / src / main"
and you would not be able to access "templates" nor "templates/home.html"

if you were in "go_web_app" and you ran "go run main.go"
then the root of your files would be "go_web_app"
and you would be able to access "templates" and "templates/home.html"

this is a good article - though advanced (so don't freak out if you're just getting started)

this can also be useful in debugging:
-- import "os"
---- log.Println("ENV: ", os.Environ())
*/
