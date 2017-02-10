package tricks

import (
	"fmt"
	"testing"
)

func TestCombination(t *testing.T) {
	combc, _ := Combination(64, 2)
	for comb := range combc {
		fmt.Println(comb)
	}
}
