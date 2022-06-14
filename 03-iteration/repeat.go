package iteration

// Repeat takes a character and returns it repeated passed number of times
func Repeat(c string, r int) string {
	var repeated string

	for i := 0; i < r; i++ {
		repeated = repeated + c
	}

	return repeated
}
