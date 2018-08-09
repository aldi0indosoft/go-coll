package main

import (
	"errors"
	"fmt"
)

// TestObject just testing
type TestObject struct {
	core int
}

// SplitHalf testing
func (n *TestObject) SplitHalf() (int, error) {

	// return 1
	if n.core%2 != 0 {
		err := errors.New("error testing")
		return 2, err
	}
	return n.core, nil
}

func main() {
	n := TestObject{core: 19}
	// se := 'test'
	c, e := n.SplitHalf()
	if e != nil {
		fmt.Print(e)
	} else {
		fmt.Print(c)
	}
	// if e != nil {
	// 	return e
	// }
}
