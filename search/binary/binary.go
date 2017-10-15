package binary

import "log"

func Search(array []int, element int) int {
	var (
		min  = 0
		max  = len(array) - 1
		curr = (max - min) / 2
	)

	// []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	// 12
	// 0 11 - i5 - v6
	// 5 11 - i8 - v9
	// 8 11 - i9 - v10
	// 9 11 - i10 - v11
	// 10 11-

	log.Printf("min: %d max: %d curr: %d", min, max, curr)

	for {
		log.Printf("min: %d max: %d curr: %d", min, max, curr)
		if array[curr] == element {
			return curr
		}

		switch {
		case array[curr] < element:
			min = curr
		case array[curr] > element:
			max = curr
		}

		curr = ((max - min) / 2) + min
	}

	return -1
}
