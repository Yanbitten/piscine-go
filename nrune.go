package piscine

func NRune(s string, n int) rune {
	result := []rune(s) // кастинг - преобразование типов. перевод строки в  массив рун

	for index, value := range result { // перевод массив рун в символы (разбиваем на отдельные элементы )
		if index == n-1 {
			return value
		}
	}
	return 0
}
