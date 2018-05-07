package main

import (
	"fmt"
	"github.com/sobamchan/project-euler/libs"
	"strconv"
)

func main() {

	var x, y int
	var z int
	var largestProduct int

	for x = 0; x < 1000; x++ {
		for y = 0; y < 1000; y++ {
			z = x * y
			zstr := strconv.Itoa(z)
			if rzstr := libs.Reverse(zstr); zstr == rzstr {
				if largestProduct < (x * y) {
					largestProduct = x * y
				}
			}
		}
	}

	fmt.Printf("%d\n", largestProduct)
}
