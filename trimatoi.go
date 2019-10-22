package piscine

func TrimAtoi(s string) int {

	k := 1    //
	yana := 0 //  мой ответ

	for _, a := range s {
		if a >= '0' && a <= '9' { // проверка на наличии цифры
			b := int(a - 48)   // если цифр, сохраняем в б
			yana = yana*10 + b // обновление положительного результата (всегда)
		} else if a == '-' && yana == 0 {
			k = -1 // при отрицательных значениях * 1
		}
	}
	return yana * k
}
