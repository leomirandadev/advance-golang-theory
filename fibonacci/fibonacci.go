package fibonacci

func RecursionWithoutMemo(n int) int {
	if n <= 1 {
		return n
	}
	return RecursionWithoutMemo(n-1) + RecursionWithoutMemo(n-2)
}
