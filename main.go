package main

import (
	"flag"
	"log"
	"os"

	"github.com/zhuliminl/easyrn-server/config"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		log.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
}
