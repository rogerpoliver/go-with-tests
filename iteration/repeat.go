package iteration

func Repeat(char string, numberOfLoops int) string {
	var repetitions string
	for i := 0; i < numberOfLoops; i++ {
		repetitions += char
	}
	return repetitions
}
