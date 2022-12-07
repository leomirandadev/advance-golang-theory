package pointer

// -------------------------------------
// -------------TASK 0------------------
// -------------------------------------
func Task0() int {
	value := 10

	changeValueInt(value)

	return value // 10
}

func changeValueInt(value int) {
	value = 30
}
