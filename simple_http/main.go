package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	r, err := http.Get("http://archive.org/wayback/available?url=beachorganicsskincare.com/Atlantis-Soap.html")
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()
	if dump, err := ioutil.ReadAll(r.Body); err == nil {
		fmt.Println(string(dump))
	}
}
