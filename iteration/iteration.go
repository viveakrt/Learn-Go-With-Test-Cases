package iteration

// Take one string and one integer input and return repeated string by integer
func Repeat(character string, iterate int) (repeated string) {

	for i := 0; i < iterate; i++ {
		repeated += character
	}
	return repeated
}
