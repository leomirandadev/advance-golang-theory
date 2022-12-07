package pointer

// -------------------------------------
// -------------TASK 1------------------
// -------------------------------------
func Task1() int {
	value := 10

	changeValuePointer(&value)

	return value // 30
}

func changeValuePointer(pointerValue *int) {
	newValue := 30
	*pointerValue = newValue
}
