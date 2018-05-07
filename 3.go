package main

import (
	"fmt"
	"github.com/sobamchan/project-euler/libs"
)

func GetPrimeFactor(x int) int {
	p := 3
	for {
		if amari := x % p; amari == 0 {
			return p
		}
		for {
			p += 1
			if libs.IsPrime(p) {
				break
			}
		}
	}
}

func main() {
	x := 600851475143

	for i := 0; i < x; i++ {
		factor := GetPrimeFactor(x)
		x = x / factor
		fmt.Printf("%d\n", factor)
	}
}
