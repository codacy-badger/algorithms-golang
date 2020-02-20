package generator

import (
	"fmt"
	"math/rand"
	"time"
)

// Generate a numerical slice
func NumericalSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())
	var slice = make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}

	fmt.Printf("Generate numerical slice: %v\n", slice)

	return slice
}
