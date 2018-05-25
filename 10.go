package main

import (
	"fmt"
	"github.com/sobamchan/project-euler/libs"
)

func main() {
	i := 2
	sum := 0

	for i < 2000000 {
		if libs.IsPrime(i) {
			sum += i
		}
		i++
	}

	fmt.Printf("%d\n", sum)
}
