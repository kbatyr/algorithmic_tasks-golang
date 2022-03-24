package student

import (
	"github.com/01-edu/z01"
)

func Output(smth []rune) {
	for _, num := range smth {
		z01.PrintRune(num)
	}
}

func Recursion(index, n int, res []rune) {
	if index == n {
		Output(res)
		if res[0] != rune(58-n) {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		} else {
			z01.PrintRune('\n')
		}
		return
	}

	if index == 0 {
		for i := '0'; i <= '9'; i++ {
			res[index] = i
			Recursion(index+1, n, res)
		}
	} else {
		for i := res[index-1] + 1; i <= '9'; i++ {
			res[index] = i
			Recursion(index+1, n, res)
		}
	}
}

func PrintCombN(n int) {
	if n < 0 || n > 9 {
		return
	}

	Recursion(0, n, make([]rune, n))
}
