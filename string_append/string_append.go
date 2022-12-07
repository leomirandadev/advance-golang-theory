package string_append

func AppendString(times int) string {
	var data string = "dal"

	for i := 0; i < times; i++ {
		data += "e"
	}

	return data
}
