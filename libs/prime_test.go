package libs

import (
	"fmt"
	"testing"
)

func TestIsPrime(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		fmt.Println(IsPrime(i))
	}
}
