package main

import (
	"bytes"
	"io/ioutil"
)

func main() {
	var dirs bytes.Buffer
	files, err := ioutil.ReadDir("./")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			// fmt.Println("Name: ", file.Name())
			dirs.WriteString(file.Name())
			dirs.WriteString("\r\n")
		}
	}
	if err := ioutil.WriteFile("list.txt", dirs.Bytes(), 0644); err != nil {
		panic(err)
	}
}
