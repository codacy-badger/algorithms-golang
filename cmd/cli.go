package main

import (
	"fmt"
	"github.com/molevs/algorithms-golang/sort/exchange/bubble"
	"github.com/molevs/algorithms-golang/sort/exchange/shaker"
	"github.com/molevs/algorithms-golang/sort/generator"
)

func main() {
	var (
		size  = 5
		slice []int
		sort  []int
	)

	slice = generator.NumericalSlice(size)

	sort = bubble.Sort(slice)
	fmt.Printf("unsort: %v\n", slice)
	fmt.Printf("sort: %v\n", sort)

	slice = generator.NumericalSlice(size)
	sort = shaker.Sort(slice)
	fmt.Printf("unsort: %v\n", slice)
	fmt.Printf("sort: %v\n", sort)
}
