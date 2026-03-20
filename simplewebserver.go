package main

import (
	"errors"
	"log"
	"net/http"
	"os"
)

func main() {
	tlsCertFile := os.Args[2]
	tlsKeyFile := os.Args[3]
	fileServerRootDir := os.Args[4]

	switch os.Args[1] {

	case "public":
		log.Fatal(http.ListenAndServeTLS(
			":8080", tlsCertFile, tlsKeyFile,
			http.FileServer(http.Dir(fileServerRootDir))))

	case "private":
		log.Fatal(http.ListenAndServeTLS(
			"127.0.0.1:8080", tlsCertFile, tlsKeyFile,
			http.FileServer(http.Dir(fileServerRootDir))))

	default:
		log.Fatalln(errors.New("invalid arguments"))

	}
}
