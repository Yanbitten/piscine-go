package piscine

func ToLower(s string) string {
	s1 := []rune(s)
	for index, letter := range s1 {
		if letter >= 'A' && letter <= 'Z' {
			s1[index] = s1[index] + 32
		}
	}
	s2 := string(s1)
	return s2
}

/*package piscine

func ToLower(s string) string {

	c := []rune(s)

	for i, a := range s {

		if a >= 'A' && a <= 'Z' {

			c[i] = a + 32

		}
	}
	return string(c)
}
*/
