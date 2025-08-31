package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gboulant/dingo-env"
)

var serverhost string
var serverport int

func main() {

	env.StringVar(&serverhost, "SERVER_HOST", "localhost", "the server host name")
	env.IntVar(&serverport, "SERVER_PORT", 1234, "the server port number")
	err := env.Parse()
	if err != nil {
		log.Printf("warning: %s", err)
	}

	flag.StringVar(&serverhost, "host", serverhost, "the server host name")
	flag.IntVar(&serverport, "port", serverport, "the server port number")
	flag.Parse()

	fmt.Printf("host: %s\n", serverhost)
	fmt.Printf("port: %d\n", serverport)

}
