// 예제 10부터 16까지

package main

import (
	//"text/template"
	"html/template" // injection proof
	"net/http"
)

// http://localhost:8080
// http://localhost:8080/mike

func main() {
	templateServer1()
}

func templateServer1() {
	http.HandleFunc("/", myHandlerFunc)
	http.HandleFunc("/mike", mikeFunc) // path 대소문자 구문하는 것 주의
	http.ListenAndServe(":8080", nil)
}

func myHandlerFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("anyNameForTemplate").Parse(docSimple)
	if err == nil {
		// tmpl.Execute(w, nil) // nil means no data to pass in
		// tmpl.Execute(w, req.URL.Path)
		tmpl.Execute(w, req.URL.Path[1:]) // 슬레시 제거
	}
}

func mikeFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")

	templates := template.New("template")
	templates.New("test").Parse(doc)
	templates.New("header").Parse(head)
	templates.New("footer").Parse(foot)

	context := Context{
		"Samil Lee",
		"more beer, please",
		req.URL.Path,
		[]string{"New Korean", "Draft", "Larger"},
		"Favorite Beers",
		"<script>alert('you have been pwned, BIATCH')</script>",
	}

	templates.Lookup("test").Execute(w, context)
}

type Context struct {
	FirstName     string
	Message       string
	URL           string
	Beers         []string
	Title         string
	ScriptMessage string
}

const docSimple = `
<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title>{{.}}</title>
</head>
<body>
    <h1>{{.}}</h1>
    {{if eq . "Lee"}}
        <h2>We're out of beer, {{.}}. Sorry!</h2>
    {{else}}
        <h2>You {{.}} not allowed</h2>
    {{end}}

    <hr></body>
</html>
`

const doc = `
{{template "header" .Title}}
<body>
    <h1>{{.FirstName}} says, "{{.Message}}"</h1>
    {{if eq .URL "/nobeer"}}
        <h2>We're out of beer, {{.FirstName}}. Sorry!</h2>
    {{else}}
        <h2>Yes, grab another beer, {{.FirstName}}</h2>
        <ul>
            {{range .Beers}}
            <li>{{.}}</li>
            {{end}}
        </ul>
    {{end}}

    <hr></body>
{{template "footer" .ScriptMessage}}
`

const head = `
<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title>{{.}}</title>
</head>
`

const foot = `
    <script>{{.}}</script>
</html>
`

/*
template.New
go doesn't have constructors like in object oriented languages
go is type based
go's struct type often used like an object
-- holds data
-- can have methods (the struct must be a "receiver" on a function)
if any initialization is needed, then create a method that does this and returns the instance:
func New(name string) *Template
New allocates a new template with the given name.
http://golang.org/pkg/text/template

template.Parse
Parse parses a string into a template.

template.Execute
merge your template with data

FUNCTION vs. METHOD


create a template that contains all of your templates
any sub-templates invoked have to be either
-- siblings
-- descendents
of the parent template

{{template "header"}}
"header" is the name we gave the template with template.New

func (*Template) Lookup
func (t *Template) Lookup(name string) *Template
Lookup returns the template with the given name that is associated with t,
or nil if there is no such template.


html/template
Package template (html/template) implements data-driven templates for generating
HTML output safe against code injection. It provides the same interface as package
text/template and should be used instead of text/template whenever the output is HTML.

*/
