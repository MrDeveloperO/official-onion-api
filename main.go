package main

import (
	"flag"
	"fmt"
)

var list = flag.String("list", "official_links.csv", "Path for links CSV file")
var addr = flag.String("addr", "localhost:8080", "Address to listen for HTTP API")

func main() {
	flag.Parse()

	onions, err := NewOnions(*list)
	if err != nil {
		panic(err)
	}

	srv := NewServer(*addr)
	srv.Register(NewController(onions))

	fmt.Println("Server starting on ", *addr)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
