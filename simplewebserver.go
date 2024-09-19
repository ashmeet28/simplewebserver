package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Fatal(http.ListenAndServe(os.Args[1]+":8080", http.FileServer(http.Dir(os.Args[2]))))
}
