package piscine

import "github.com/01-edu/z01"

func PrintAlphabet() {

	for i := 97; i < 123; i++ {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')  

	
}
