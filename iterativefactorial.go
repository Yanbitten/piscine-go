package piscine

func IterativeFactorial(nb int) int {

	if nb < 0 || nb > 12 { // если исходное значение - отрицательное или охереть какое больлшое
		return 0 //то функция выводит 0
	} else if nb == 0 { //если же исходное значение равно 0
		return 1 //то функция выводит 0
	}

	c := nb                      //создаем переменную и приравниеваем ее к исходному значению
	b := 1                       //создаем новую переменную и приравниваем ее к 1
	for nb := 1; nb <= c; nb++ { //пишем цикл при котором конечное значение цикла ровно исходному значению

		b = b * nb //производим функциональное вычисление
	}
	return b //функция возвращается в b

}
