package main

import (
	"fmt"
)

func is_dividable(x, a int) bool {
	r := x % a
	if r == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	ar := []int{1, 2}
	for {
		if x := ar[len(ar)-1] + ar[len(ar)-2]; x < 4000000 {
			ar = append(ar, x)
		} else {
			break
		}
	}
	var sum int
	for _, v := range ar {
		if is_dividable(v, 2) {
			sum += v
		}
	}
	fmt.Printf("%d\n", sum)
}
