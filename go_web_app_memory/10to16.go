// 예제 10부터 16까지

package main

import (
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
		tmpl.Execute(w, req.URL.Path[1:])
	}
}

func mikeFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	//tmpl, err := template.New("anyNameForTemplate").Parse(doc)
	templates := template.New("template")
	templates.New("test").Parse(doc)
	templates.New("header").Parse(head)
	templates.New("footer").Parse(foot)

	//if err == nil {
	context := Context{
		"Samil Lee",
		"more beer, please",
		req.URL.Path,
		[]string{"New Korean", "Draft", "Larger"},
		"Favorite Beers",
		"<script>alert('you have been pwned, BIATCH')</script>",
	}
	// tmpl.Execute(w, context)
	templates.Lookup("test").Execute(w, context)
	//}
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
my description: create the template "object"
go doesn't have constructors like in object oriented languages
go is type based
go's struct type often used like an object
-- holds data
-- can have methods (the struct must be a "receiver" on a function)
if any initialization is needed, then create a method that does this and returns the instance:
func New(name string) *Template
New allocates a new template with the given name.
http://golang.org/pkg/text/template/#New

template.Parse
my description: put your template into the template "object"
func (t *Template) Parse(text string) (*Template, error)
Parse parses a string into a template. Nested template definitions will be associated
with the top-level template t. Parse may be called multiple times to parse definitions
of templates to associate with t.
http://golang.org/pkg/text/template/#Template.Parse

template.Execute
my description: merge your template with data
func (t *Template) Execute(wr io.Writer, data interface{}) (err error)
Execute applies a parsed template to the specified data object, and writes the output to wr.
If an error occurs executing the template or writing its output, execution stops,
but partial results may already have been written to the output writer.
A template may be executed safely in parallel.
http://golang.org/pkg/text/template/#Template.Execute



FUNCTION
func New(name string) *Template
New allocates a new template with the given name.
http://golang.org/pkg/text/template/#New

METHOD
func (t *Template) New(name string) *Template
New allocates a new template associated with the given one and with the same delimiters.
The association, which is transitive, allows one template to invoke another with a {{template}} action.
http://golang.org/pkg/text/template/#Template.New

*/

/*
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
*/

/*
html/template
Package template (html/template) implements data-driven templates for generating
HTML output safe against code injection. It provides the same interface as package
text/template and should be used instead of text/template whenever the output is HTML.

HTML templates treat data values as plain text which should be encoded so they can be
safely embedded in an HTML document. The escaping is contextual, so actions can appear
within JavaScript, CSS, and URI contexts.

http://golang.org/pkg/html/template/

to run the above code ...
try the different Message variables with text/template import
... then ...
try the different Message variables with html/template import
*/
