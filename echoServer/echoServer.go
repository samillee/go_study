package main

import (
	//"io"
	"fmt"
	"log"
	//"net"
	"net/http"
)

const listenAddr = "localhost:4231"

func main() {
	/*
		l, err := net.Listen("tcp", listenAddr)
		if err != nil {
			log.Fatal(err)
		}
		for {
			c, err := l.Accept()
			if err != nil {
				log.Fatal(err)
			}
			//io.Copy(c, c)
			go io.Copy(c, c)
		}
	*/
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Mike, web")
}
