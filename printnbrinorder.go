package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {

		z01.PrintRune('0')
	}
	array := []int{} //massiv

	for n != 0 { //n не равно нулю
		balance := n % 10              // новые элементы для массива
		array = append(array, balance) // сохраняем элементы старого + нового

		n = n / 10 // убираем осток от числа
	}

	// fmt.Println(array)

	len := 0          // вычисление длинны
	for range array { // пройти по всей длинне массива
		len++
	}

	// bubble sort
	for k := 1; k < len; k++ { // создаем цикл который должен пройти от начала до конца len раз
		for i := 1; i < len; i++ { // дабл цикл : (если мы пройдем один раз, сортировка не сработает
			if array[i-1] > array[i] { // проверка если левое > правого то
				array[i-1], array[i] = array[i], array[i-1] // замена левого на правый
			}
		}
	}

	// output
	for _, nb := range array { // проходиться по всему масиву
		z01.PrintRune(rune(nb + 48)) //  оборачиваем в руну это называется каст
	}

}
