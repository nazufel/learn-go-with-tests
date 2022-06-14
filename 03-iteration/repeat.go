package iteration

const repeatCount = 5

// Repeat takes a character and returns it repeated
func Repeat(c string) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated = repeated + c
	}

	return repeated
}
