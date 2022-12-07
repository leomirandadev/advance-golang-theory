package scheduler

func Task2(times, multiplier int) []int {
	var data []int = make([]int, 0, times)

	for i := 0; i < times; i++ {
		data = append(data, i*multiplier)
	}

	return data
}
