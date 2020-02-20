// https://en.wikipedia.org/wiki/Bubble_sort
package bubble

// Bubble sort
func Sort(input []int) []int {
	length := len(input)

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if input[i] < input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}

	return input
}
