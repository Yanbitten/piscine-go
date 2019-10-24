package piscine

func Capitalize(s string) string {

	sring := []rune(s) // кастинг : создаем массив рун cast

	len := 0 // вычисление длины строки
	for range sring {
		len++
	}
	for i, bykva := range sring { // дайет доступ к каждой букве и диджителу

		if i == 0 || !isAlphaNum(sring[i-1]) { // проверяем если первая буква в слове

			if bykva >= 'a' && bykva <= 'z' { // проверяем если маленькая
				sring[i] = bykva - 32 // замена на большую
			}
		} else {
			if bykva >= 'A' && bykva <= 'Z' { // если это большая
				sring[i] = bykva + 32 // замена на маленькую
			}
		}

	}

	return string(sring) // возвращаем  тип стринг

}

func isAlphaNum(a rune) bool { // сделали чтобы проверить

	if a >= 'a' && a <= 'z' { // маленькая ли буква
		return true
	}
	if a >= 'A' && a <= 'Z' { // большая ли буква
		return true
	}
	if a >= '0' && a <= '9' { // цифра ли
		return true
	}

	return false
}
