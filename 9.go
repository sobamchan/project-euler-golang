package main

import "fmt"

func main() {
	target := 1000
	var c int

	for a := 1; a < target+1; a++ {
		for b := 1; b < target+1; b++ {
			c = target - a - b
			if c*c == a*a+b*b {
				fmt.Printf("a: %d, b: %d, c: %d\n", a, b, c)
				fmt.Printf("a * b * c: %d\n", a*b*c)
			}
		}
	}
}
