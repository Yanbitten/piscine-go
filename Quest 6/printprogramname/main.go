package main

import (
	"fmt"
	"os"
)

func main() {
	// runes := []rune(os.Args[0])    // casting - преобразование типов . перевод типов. Мы переводим строку в массив Рун,
	// for _, letter := range runes { // тк индекс нам не нужен мы его пропускаем (for _,) затем переводим массив Рун в символы
	// 	z01.PrintRune(letter)

	// }
	// z01.PrintRune('\n')
	arguments := os.Args
	fmt.Println(arguments[0])
}
