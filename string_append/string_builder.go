package string_append

import "strings"

func AppendStringBuilder(times int) string {
	var str strings.Builder

	str.WriteString("dal")

	for i := 0; i < times; i++ {
		str.WriteString("e")
	}

	return str.String()
}
