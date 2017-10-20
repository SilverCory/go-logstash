package main

import (
	"flag"
	"github.com/SilverCory/go-logstash/http"
	"github.com/SilverCory/go-logstash/log"
)

func main() {
	authKeyPtr := flag.String("authkey", "", "The authorisation key. Make this long enought it will take a while to guess.")
	flag.Parse()

	s := http.New(*authKeyPtr, log.New())
	s.Open()
}
