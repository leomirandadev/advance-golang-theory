package pointer

import (
	"fmt"
)

// -------------------------------------
// -------------TASK 3------------------
// -------------------------------------
func Task3() []int {
	value := []int{1, 2, 3}

	changeSizeSlice(value)

	return value // {1, 2, 3}
}

// head slice -> len, cap
// array
// [__|__|__]

func changeSizeSlice(slice []int) {
	slice = slice[:1]
	fmt.Println(len(slice), cap(slice)) // len 1 cap 3
}
