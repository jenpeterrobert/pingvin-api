package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/jenpeterrobert/pingvin-api/pkg/handlers"
)

const (
	defaultAddress string = "0.0.0.0"
	defaultPort    int    = 80
)

var (
	address string
	port    int
)

func main() {

	flag.StringVar(&address, "addr", defaultAddress, "Server address")
	flag.IntVar(&port, "port", defaultPort, "Server port")
	flag.Parse()
	runAPI()
}

func runAPI() {

	fullAddress := fmt.Sprintf("%s:%d", address, port)
	http.HandleFunc("/test", handlers.TestHandler)
	http.ListenAndServe(fullAddress, nil)
}
