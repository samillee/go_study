package main

import (
	"bufio"
	"github.com/samillee/go_study/go_web_app_mvc/viewmodels"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

// http://localhost:8080/home
func main() {

	http.Handle("/", new(MyHandler))
	//http.HandleFunc("/", new(MyHandler)) // 함수인지 핸들러 형인지 주의 해야함
	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/scripts/", serveResource)
	http.HandleFunc("/video/", serveResource)
	http.ListenAndServe(":8080", nil)
}

type MyHandler struct{}

func (this MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	templates := populateTemplates()

	requestedFile := req.URL.Path[1:]
	template := templates.Lookup(requestedFile + ".html")

	var context interface{} = nil

	switch requestedFile {
	case "home":
		context = viewmodels.GetHome()
	case "search":
		context = viewmodels.GetSearch()
	}

	if template != nil {
		template.Execute(w, context)
	} else {
		w.WriteHeader(404)
	}
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpg"
	} else if strings.HasSuffix(path, ".svg") {
		contentType = "image/svg+xml"
	} else if strings.HasSuffix(path, ".mp4") {
		contentType = "video/mp4"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else {
		contentType = "text/plain"
	}

	log.Println(path)
	log.Println(contentType)

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}

func populateTemplates() *template.Template {
	result := template.New("templates")

	// 특정 폴더에서 html 파일 찾도록 함
	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)
	// -1 means all of the contents
	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		log.Println(pathInfo.Name())
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths,
				basePath+"/"+pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)

	return result
}

/*
MODEL
business logic & rules
data storage

VIEW
what the client sees

CONTROLLER
the glue between model & view
coordinates the model & view layers
determines how the model needs to be interacted with to meet a user's request
passes the results of the model layers work to the view layer
responsibilities:
- generate output and send it back to client
-- templates
-- bind data
- receive user actions
-- ajax
-- forms

change your references in your HTML files from this stuff ...
<link rel='stylesheet' href='../public/css/flyout_menu.css'>
... to this stuff ...
<link rel='stylesheet' href='/css/flyout_menu.css'>

test it here
http://localhost:8080/home
*/

/*
we're going to break our main html page down into different parts
--- header
--- content1
--- content2
--- footer

This will help with
-- code reusability
-- organizing our data and keeping it clean

we are going to separate the data that is used in the VIEW layer
from the rest of the data that the application uses
-- good practice as the needs of the VIEW and MODEL layer differ over time

create:
viewmodels / home.go

add this to main.go imports:
"viewmodels"

test it here
http://localhost:8080/home


MODEL
business logic & rules
data storage

VIEW
what the client sees

CONTROLLER
the glue between model & view
coordinates the model & view layers
determines how the model needs to be interacted with to meet a user's request
passes the results of the model layers work to the view layer
responsibilities:
- generate output and send it back to client
-- templates
-- bind data
- receive user actions
-- ajax
-- forms
*/
