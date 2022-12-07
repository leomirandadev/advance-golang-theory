package scheduler

func Task3(times, multiplier int) []int {
	var data []int = make([]int, times)

	for i := 0; i < times; i++ {
		data[i] = i * multiplier
	}

	return data
}
