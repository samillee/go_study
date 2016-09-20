// https://github.com/GoesToEleven/GolangTraining
// GolangTraining\27_code-in-process\42_Http\
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

//http://localhost:8000
func main() {
	//readHttpHeaderServer(8000)
	simpleHttpServer(8001)
}

func simpleHttpServer(port int) {
	server, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handleSimpleHttpConn(conn)
	}
}

func handleSimpleHttpConn(conn net.Conn) {
	defer conn.Close()

	// for header processing
	var url, method, hString string
	headers := map[string]string{}
	scanner := bufio.NewScanner(conn)
	i := 0
	for scanner.Scan() {
		ln := scanner.Text()
		//fmt.Println(ln) // 서비스 중에는 프린트 안돼네
		hString = hString + fmt.Sprintln(ln)

		if i == 0 {
			fs := strings.Fields(ln)
			method = fs[0]
			url = fs[1]
			//fmt.Println("METHOD", method)
		} else {
			// in headers now
			// when line is empty, header is done
			if ln == "" {
				break
			}
			fs := strings.SplitN(ln, ": ", 2)
			headers[fs[0]] = fs[1]
		}

		i++
	}

	// response
	body := fmt.Sprintf("%s %s \n\n%v\n%v\n%v", method, url, hString, "hello world 2", headers)

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handleSimpleHttpConn2(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	i := 0
	headers := map[string]string{}
	//var url, method string
	var method string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			fs := strings.Fields(ln)
			method = fs[0]
			//url = fs[1]
			//fmt.Println("METHOD", method)
			//fmt.Println("URL", url)
		} else {
			// in headers now
			// when line is empty, header is done
			if ln == "" {
				break
			}
			fs := strings.SplitN(ln, ": ", 2)
			headers[fs[0]] = fs[1]
		}

		i++
	}

	// parse body
	if method == "POST" || method == "PUT" {
		amt, _ := strconv.Atoi(headers["Content-Length"])
		buf := make([]byte, amt)
		io.ReadFull(conn, buf)
		// 주의 변경해 줘야 함 in buf we will have the POST content
		fmt.Println("BODY:", string(buf))
	}

	// response
	body := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
	<form method="POST">
		<input type="text" name="key" value="">
		<input type="submit">
	</form>
</body>
</html>
	`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
