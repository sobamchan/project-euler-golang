package main

import (
	"fmt"
	"github.com/sobamchan/project-euler/libs"
)

func main() {
	i := 2
	sum := 0

	for {
		if libs.IsPrime(i) {
			sum += i
		}
		i++
		if i > 2e+10 {
			break
		}
	}

	fmt.Printf("%d\n", sum)
}
