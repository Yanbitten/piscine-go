package piscine

func ToUpper(s string) string {
	s1 := []rune(s)
	for index, letter := range s1 {
		if letter >= 'a' && letter <= 'z' {
			s1[index] = s1[index] - 32
		}
	}
	s2 := string(s1)
	return s2
}

/*package piscine

func ToUpper(s string) string {

	c := []rune(s)

	for i, a := range s {

		if a >= 'a' && a <= 'z' {

			c[i] = a - 32

		}
	}
	return string(c)
}
*/
