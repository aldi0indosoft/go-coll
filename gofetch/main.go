package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://google.co.id", nil)
	if err != nil {
		fmt.Printf("Response %s \n", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Response %s \n", err)
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Response %s \n", err)
	}

	fmt.Printf("Response %s \n", []byte(html))
}
