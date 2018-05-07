package main

import "fmt"

func isOK(x int) bool {
	for i := 1; i < 21; i++ {
		if (x % i) != 0 {
			return false
		}
	}
	return true
}

func main() {
	x := 1
	for {
		if isOK(x) {
			fmt.Printf("%d \n", x)
			break
		}
		x++
	}
}
