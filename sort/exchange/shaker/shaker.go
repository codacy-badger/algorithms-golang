// https://en.wikipedia.org/wiki/Cocktail_shaker_sort
package shaker

// Cocktail shaker sort
func Sort(input []int) []int {
	var (
		left  int
		right = len(input) -1
	)

	for left < right {
		for i := left; i < right; i++ {
			if input[i] > input[i+1] {
				input[i+1], input[i] = input[i], input[i+1]
			}
		}

		right--

		for j := right; j > left; j-- {
			if input[j-1] > input[j] {
				input[j], input[j-1] = input[j-1], input[j]
			}
		}

		left++
	}

	return input
}
