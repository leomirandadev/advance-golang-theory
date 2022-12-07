package pointer

import "testing"

func TestTask0(t *testing.T) {
	result := Task0()
	if result != 10 {
		t.Error("Error task1")
	}

}

func TestTask1(t *testing.T) {
	result := Task1()
	if result != 30 {
		t.Error("Error task1")
	}
}

func TestTask2(t *testing.T) {
	result := Task2()
	if !isEqualSlice(result, []int{1, 30, 3}) {
		t.Error("Error task2")
	}
}

func TestTask3(t *testing.T) {
	result := Task3()
	if !isEqualSlice(result, []int{1, 2, 3}) {
		t.Error("Error task2")
	}
}

func TestTask4(t *testing.T) {
	result := Task4()
	if !isEqualArray(result, [3]int{1, 2, 3}) {
		t.Error("Error task2")
	}
}

func isEqualArray(a, b [3]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func isEqualSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
