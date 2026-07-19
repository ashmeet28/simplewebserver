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
	fileServerRootDirPrefix := os.Args[5]

	switch os.Args[1] {

	case "public":
		http.Handle(fileServerRootDirPrefix, http.StripPrefix(
			fileServerRootDirPrefix, http.FileServer(http.Dir(fileServerRootDir))))
		err := http.ListenAndServeTLS(":8080", tlsCertFile, tlsKeyFile, nil)
		log.Fatal(err)

	case "private":
		http.Handle(fileServerRootDirPrefix, http.StripPrefix(
			fileServerRootDirPrefix, http.FileServer(http.Dir(fileServerRootDir))))
		err := http.ListenAndServeTLS("127.0.0.1:8080", tlsCertFile, tlsKeyFile, nil)
		log.Fatal(err)

	default:
		log.Fatalln(errors.New("invalid arguments"))

	}
}
