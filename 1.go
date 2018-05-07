package main

import "fmt"

func is_dividable(x, a int) bool {
	r := x % a
	if r == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	sum := 0
	for i := 1; i < 1000; i++ {
		if is_dividable(i, 3) || is_dividable(i, 5) {
			sum += i
		}
	}
	fmt.Printf("%d\n", sum)
}
