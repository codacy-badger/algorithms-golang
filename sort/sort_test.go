package sort

import (
	"fmt"
	"github.com/molevs/algorithms-golang/sort/exchange/bubble"
	"github.com/molevs/algorithms-golang/sort/exchange/shaker"
	"github.com/molevs/algorithms-golang/sort/generator"
	"testing"
)

func TestBubble(t *testing.T) {
	testSort(bubble.Sort(generator.NumericalSlice(5)), t)
}

func TestShaker(t *testing.T) {
	testSort(shaker.Sort(generator.NumericalSlice(5)), t)
}

func testSort(sorted []int, t *testing.T) {
	length := len(sorted) - 1
	order := ">"

	for i := 0; i < length; i++ {
		if sorted[i] > sorted[i+1] {
			t.Error("Сортировочка не удалась", sorted[i], order, sorted[i+1])
		} else {
			fmt.Printf("%v %s %v\n", sorted[i], order, sorted[i+1])
		}
	}
}
