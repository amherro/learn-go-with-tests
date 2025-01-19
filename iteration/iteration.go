package iteration

func Iterate(character string, repeatCount int) string {
	var charString string
	for i := 0; i < repeatCount; i++ {
		charString += character
	}
	return charString
}
