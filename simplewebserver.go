package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	if os.Args[1] == "public" {
		log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(os.Args[2]))))
	} else if os.Args[1] == "private" {
		log.Fatal(http.ListenAndServe("127.0.0.1:8080", http.FileServer(http.Dir(os.Args[2]))))
	}
}
