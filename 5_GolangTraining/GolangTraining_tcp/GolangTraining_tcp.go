// -----------------------------------------------
// https://github.com/GoesToEleven/GolangTraining
// 소스에서 Json 관련 참고할 만한 것 추출

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	//simpleListener(9000)
	//simpleDial(9000)
	//simpleEchoServer(9001)
	RadisServer(5231)
}

func simpleListener(port int) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		io.WriteString(conn, fmt.Sprint("Hello World\n", time.Now(), "\n"))

		conn.Close()
	}
}

func simpleDial(port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bs, _ := ioutil.ReadAll(conn)
	fmt.Println(string(bs))
}

func simpleEchoServer(port int) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		// handles unlimited connections
		go func() {
			io.Copy(conn, conn)
			conn.Close()
		}()
	}
}

type Command struct {
	Fields []string
	Result chan string
}

func redisHandle(commands chan Command) {
	var data = make(map[string]string)
	for cmd := range commands {
		if len(cmd.Fields) < 2 {
			cmd.Result <- "Expected at least 2 arguments"
			continue
		}

		fmt.Println("GOT COMMAND", cmd)

		switch cmd.Fields[0] {
		// GET <KEY>
		case "GET":
			key := cmd.Fields[1]
			value := data[key]

			cmd.Result <- value

		// SET <KEY> <VALUE>
		case "SET":
			if len(cmd.Fields) != 3 {
				cmd.Result <- "EXPECTED VALUE"
				continue
			}
			key := cmd.Fields[1]
			value := cmd.Fields[2]
			data[key] = value
			cmd.Result <- ""
		// DEL <KEY>
		case "DEL":
			key := cmd.Fields[1]
			delete(data, key)
			cmd.Result <- ""
		default:
			cmd.Result <- "INVALID COMMAND " + cmd.Fields[0] + "\n"
		}
	}
}

func connHandle(commands chan Command, conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		result := make(chan string)
		commands <- Command{
			Fields: fs,
			Result: result,
		}

		io.WriteString(conn, <-result+"\n")
	}

}

/*
you can now SET and GET and DEL

~ $ telnet localhost 9000
SET fav pop
GET fav
pop
DEL fav
GET fav
*/
func RadisServer(port int) {
	li, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	commands := make(chan Command)
	go redisHandle(commands)

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go connHandle(commands, conn)
	}
}
