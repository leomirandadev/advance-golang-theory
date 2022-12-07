package scheduler

func Task1(times, multiplier int) []int {
	var data []int

	for i := 0; i < times; i++ {
		data = append(data, i*multiplier)
	}

	return data
}
