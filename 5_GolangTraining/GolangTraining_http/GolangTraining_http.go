// https://github.com/GoesToEleven/GolangTraining
// GolangTraining\27_code-in-process\43_Http-Server ...\
package main

// go get github.com/nu7hatch/gouuid
// go get github.com/gorilla/sessions
import (
	"crypto/hmac" // 해당 장치에서의 고유키 생성하는 로직 제공
	"crypto/sha256"
	"fmt"
	"github.com/gorilla/sessions" // session 처리에 좋음, 단순 쿠키 사용 보다 이걸 사용 권장
	"github.com/nu7hatch/gouuid"  // 몇가지 UUID 생성 로직 제공
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type myHandler int

func (h myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//io.WriteString(resp, req.RequestURI)
	//io.WriteString(resp, req.URL.RequestURI())

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	switch req.URL.Path {
	case "/cat":
		io.WriteString(res, `<img src="https://upload.wikimedia.org/wikipedia/commons/0/06/Kitten_in_Rizal_Park%2C_Manila.jpg">`)
	case "/dog":
		io.WriteString(res, `<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">`)
	}

}

// ----------------------------------

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

var store = sessions.NewCookieStore([]byte("secret-password"))

func handleRoot(res http.ResponseWriter, req *http.Request) {
	/*
		// Manual session set
		cookie, err := req.Cookie("session-id")
		if err != nil {
			//id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name: "session-id",
			}
		}

		if req.FormValue("email") != "" {
			cookie.Value = req.FormValue("email")
		}

		code := getCode(cookie.Value)
		cookie.Value = code + "|" + cookie.Value

		// this doesn't run, need more code added to work
		// just shown for example of how to do auth with HMAC

		http.SetCookie(res, cookie)
		sessionValue := cookie.Value
	*/

	session, _ := store.Get(req, "session")
	if req.FormValue("email") != "" {
		// check password
		session.Values["email"] = req.FormValue("email")
	}
	session.Save(req, res)

	countCookie, err := req.Cookie("count-cookie")
	// there is no cookie
	if err == http.ErrNoCookie {
		countCookie = &http.Cookie{
			Name:  "count-cookie",
			Value: "0",
		}
	}

	// there is a cookie
	count, _ := strconv.Atoi(countCookie.Value)
	if req.Method == "GET" && req.URL.Path == "/" {
		count++
		countCookie.Value = strconv.Itoa(count)

		http.SetCookie(res, countCookie)
	}

	id, _ := uuid.NewV4() // 랜덤 UUID 생성
	machineVeriCode := getCode("machine")
	sessionValue := fmt.Sprint("UUID V4: ", id, "<br/>", "Machine code: ", machineVeriCode, "<br/>")
	sessionValue += fmt.Sprintf("%s : Count %d<br/>", session.Values["email"], count)

	//var html string
	html := `<!DOCTYPE html>
<html>
  <body>
    <form method="POST">
    ` + sessionValue + `
      <label>email</label>    
      <input type="email" name="email">
      <label>password</label>    
      <input type="password" name="password">
      <input type="submit">
    </form>
  </body>
</html>`

	io.WriteString(res, html)
}

func handleLogin(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("logged-in")
	// no cookie
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "logged-in",
			Value: "0",
		}
	}

	// check log in
	if req.Method == "POST" {
		password := req.FormValue("password")
		if password == "secret" {
			cookie = &http.Cookie{
				Name:  "logged-in",
				Value: "1",
			}
		}
	}

	// if logout, then logout
	if req.URL.Path == "/logout" {
		cookie = &http.Cookie{
			Name:   "logged-in",
			Value:  "0",
			MaxAge: -1,
		}
	}

	http.SetCookie(res, cookie)

	var html string

	// not logged in
	if cookie.Value == "0" {
		html = `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
			<h1>LOG IN</h1>
			<form method="post" action="/login">
				<h3>User name</h3>
				<input type="text" name="userName" id="userName">
				<h3>Password</h3>
				<input type="text" name="password" id="password">
				<br>
				<input type="submit">
				<!-- <input type="submit" name="logout" value="logout"> -->
			</form>
			</body>
			</html>`
	} else {
		html = `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
			<h1><a href="/logout">LOG OUT</a></h1>
			</body>
			</html>`
	}

	io.WriteString(res, html)
}

func youUp(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "catty catty catty")
}

func saveFile(res http.ResponseWriter, req *http.Request) {
	// receive form submission
	if req.Method == "POST" {
		src, _, err := req.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer src.Close()

		dst, err := os.Create(filepath.Join("./", "file.txt"))
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer dst.Close()

		io.Copy(dst, src)
	}

	// parse template
	tpl, err := template.ParseFiles("form.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute template
	err = tpl.Execute(res, nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
		log.Fatalln(err)
	}
}

func redir(w http.ResponseWriter, req *http.Request) {
	//http.Redirect(w, req, "https://localhost:7443/"+req.RequestURI, http.StatusMovedPermanently)
	http.Redirect(w, req, "https://localhost:7443"+req.RequestURI, http.StatusMovedPermanently)
}

func main() {

	/*
		var h myHandler
		http.ListenAndServe(":9000", h)
	*/

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/login/", handleLogin)  // login하고만 매칭되는 경우
	http.HandleFunc("/logout/", handleLogin) // logout하고만 매칭되는 경우
	http.HandleFunc("/cat/", youUp)
	http.HandleFunc("/upload/", saveFile)
	http.Handle("/download", http.FileServer(http.Dir("./download")))

	go http.ListenAndServe(":7000", nil)
	go http.ListenAndServe(":7001", http.RedirectHandler("https://localhost:7443/", 301))
	go http.ListenAndServe(":7002", http.RedirectHandler("https://localhost:7443/", http.StatusMovedPermanently))
	go http.ListenAndServe(":7003", http.HandlerFunc(redir))
	go http.ListenAndServe(":7011", http.RedirectHandler("https://127.0.0.1:7443/", 301))

	// openssl genrsa 1024 > key.pem  // 개인키 생성
	// openssl req -new -key key.pem > csr.pem  // 인증서 요청 생성
	// openssl req -key key.pem -x509 -nodes -sha1 -days 365 -in csr.pem -out cert.pem  // 사설 공개키 생성
	err := http.ListenAndServeTLS(":7443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
