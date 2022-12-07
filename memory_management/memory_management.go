package memory_management

func Task() {

	printInt(2)

	value2 := []int{10, 30, 50}
	printSlice(value2)

	value3 := [3]int{10, 30, 50}
	printArray(value3)

	value4 := "text"
	printString(value4)
}

// GC has 3 phases:
// - initialize agent
// - mark data to remove
// - sweep marked data
// When any of these phases are executing, the "whole world stop",
// it means that the execution of your program stop.
