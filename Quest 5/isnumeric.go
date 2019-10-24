package piscine

func IsNumeric(str string) bool {

	s := []rune(str)
	length := 0
	for range str {
		length++
	}
	for i := 0; i < length; i++ {
		if !isNumA(s[i]) {
			return false
		}
	}
	return true

}

func isNumA(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}
