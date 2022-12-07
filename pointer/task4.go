package pointer

// -------------------------------------
// -------------TASK 4------------------
// -------------------------------------
func Task4() [3]int {
	value := [3]int{1, 2, 3}

	changeArray(value)

	return value // {1, 2, 3}
}

func changeArray(slice [3]int) {
	slice[1] = 30
}
