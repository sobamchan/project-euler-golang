package main

import "fmt"

func main() {
	var sum int
	var sqrdsum int

	for i := 1; i < 101; i++ {
		sum += i
		sqrdsum += i * i
	}

	fmt.Printf("%d\n", (sum*sum)-sqrdsum)
}
