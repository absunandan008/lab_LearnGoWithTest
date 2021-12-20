package iteration

//const repeatCount = 5

// Repeates inout character for 5 times and appends it and sends it back
func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}

/*
Iteration 1
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
*/
