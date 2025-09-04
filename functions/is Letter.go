package functions

func IsLetter(text string) bool {
	valid := true
	for _, r := range text {
		if r < 32 || r > 126 {
			valid = false
			break
		}
	}
	return valid
}
