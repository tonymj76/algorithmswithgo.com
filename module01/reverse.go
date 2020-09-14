package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	torune := []rune(word)
	for i, j := 0, len(torune)-1; i < j; i, j = i+1, j-1 {
		torune[i], torune[j] = torune[j], torune[i]
	}
	return string(torune)
}
