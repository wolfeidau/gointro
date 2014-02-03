package strings

// Reverse the string passed in taking into
// consideration double width characters
func Reverse(str string) string {
	n := len(str)
	runes := make([]rune, n)
	for _, rune := range str {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}
