package main

import (
	"fmt"
	"time"
)

// TickerAt sleep execution for 2 secs
func TickerAt(tc *time.Ticker) {
	for t := range tc.C {
		fmt.Println("Tick at", t)
	}
}

func main() {
	tc := time.NewTicker(time.Second * 1)
	go TickerAt(tc)
	time.Sleep(time.Second * 5)
	tc.Stop()
	fmt.Println("Ticker Stop")
}
