package iteration

func Repeat(char string, repeatCount int) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated = repeated + char
	}

	return repeated
}
