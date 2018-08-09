package main

import (
	"flag"
	"fmt"
)

func main() {
	s := flag.String("s", "default", "type the message")
	flag.Parse()

	fmt.Println("the message:", *s)
}
