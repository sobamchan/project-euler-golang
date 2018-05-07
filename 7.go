package main

import (
	"fmt"
	"github.com/sobamchan/project-euler/libs"
)

func main() {
	c := 1
	var i int

	for c <= 10001 {
		i++
		if libs.IsPrime(i) {
			c++
		}
	}

	fmt.Printf("%d\n", i)
}
