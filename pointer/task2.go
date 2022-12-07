package pointer

// -------------------------------------
// -------------TASK 2------------------
// -------------------------------------
func Task2() []int {
	value := []int{1, 2, 3}

	newValue := value

	changeSlice(newValue)

	return value // {1, 30, 3}
}

func changeSlice(slice []int) {
	slice[1] = 30
}
