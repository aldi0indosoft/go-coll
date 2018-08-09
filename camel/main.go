package main

import (
	"camel"
	"fmt"
)

func main() {
	onta := 4
	roy := []int{
		2, 3, 1, 4,
	}
	ananto := []int{
		2, 1, 4, 3,
	}
	abi := []int{
		2, 4, 3, 1,
	}
	e, _ := camel.Tew(onta, roy, ananto, abi)
	fmt.Println("result: ", e)
}
