package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	switch os.Args[1] {
	case "public":
		log.Fatal(http.ListenAndServe(":8080",
			http.FileServer(http.Dir(os.Args[2]))))
	case "private":
		log.Fatal(http.ListenAndServe("127.0.0.1:8080",
			http.FileServer(http.Dir(os.Args[2]))))
	}
}
