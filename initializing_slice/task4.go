package scheduler

func Task4(multiplier int) [50]int {
	var data [50]int

	for i := 0; i < 50; i++ {
		data[i] = i * multiplier
	}

	return data
}
