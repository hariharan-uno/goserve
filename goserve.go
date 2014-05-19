package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port string

func init() {
	flag.StringVar(&port, "port", ":8050", "the port in which you want to serve")
}

func main() {
	flag.Parse()
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get the working directory")
	}
	fmt.Printf("Open http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(path))))
}
