package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Проверка арг-та с консоли
func isOperand(s string) bool {
	ops := [5]string{"+", "-", "*", "/", "%"}
	for _, elem := range ops {
		if elem == s {
			return true
		}
	}
	return false
}

// Печать стринга с пом. принтрун
func PrintStr(s string) {
	for _, let := range s {
		z01.PrintRune(let)
	}
	z01.PrintRune(10)
}

func atoi(s string) (int, bool) {
	var res int
	s1 := s
	if len(s) == 0 {
		return 0, false
	} else {
		if s[0] == '+' || s[0] == '-' {
			s = s[1:]
		}
		for _, ch := range s {
			if ch >= 48 && ch <= 57 {
				ch = ch - 48
				res = res*10 + int(ch)
				if res < 0 {
					break
				}
			} else {
				return 0, false
			}
		}
		if s1[0] == '-' {
			res = -res
		}
		if (s1[0] == '-' && res > 0) || (s1[0] == '+' || (s1[0] >= '0' && s1[0] <= '9')) && res < 0 {
			return 0, false
		}
	}
	return res, true
}

// конв-я числа из инт в стринг
func Itoa(nb int) (res string) {
	var negative bool
	var mod int
	var rev string
	if nb == 0 {
		return "0"
	} else if nb < 0 {
		negative = true
	}
	for nb != 0 {
		mod = nb % 10
		if negative {
			mod = -mod
		}
		rev += string(rune(mod) + 48)
		nb = nb / 10
	}
	for i := len(rev) - 1; i >= 0; i-- {
		res += string(rev[i])
	}
	if negative {
		res = "-" + res
	}
	return res
}

func main() {
	args := os.Args[1:]

	if len(os.Args) == 4 {

		// проверка на валидность значений введенных в консоль
		if isOperand(args[1]) {

			var res int

			num1, null1 := atoi(args[0])
			num2, null2 := atoi(args[2])

			// проверка чисел на валидность
			if !null1 || !null2 {
				z01.PrintRune('0')
				z01.PrintRune('\n')
				os.Exit(0)
			}

			if args[1] == "/" && num2 == 0 {
				PrintStr("No division by 0")
			} else if args[1] == "%" && num2 == 0 {
				PrintStr("No modulo by 0")
			} else if args[1] == "+" {
				res = num1 + num2
				// проверка на оверфлоу
				if (num2 >= 0 && res >= num1) || (num2 < 0 && res < num1) {
					PrintStr(Itoa(res))
				} else {
					z01.PrintRune('0')
					z01.PrintRune('\n')
				}
			} else if args[1] == "-" {
				res = num1 - num2
				// проверка на оверфлоу
				if (num2 >= 0 && res <= num1) || (num2 < 0 && res > num1) {
					PrintStr(Itoa(res))
				} else {
					z01.PrintRune('0')
					z01.PrintRune('\n')
				}
			} else if args[1] == "*" {
				res = num1 * num2
				// проверка на оверфлоу
				if res/num1 == num2 && res/num2 == num1 {
					PrintStr(Itoa(res))
				} else {
					z01.PrintRune('0')
					z01.PrintRune('\n')
				}
			} else if args[1] == "/" {
				res = num1 / num2
				// проверка на оверфлоу
				if num1 < 0 && num2 < 0 && res < 0 {
					z01.PrintRune('0')
					z01.PrintRune('\n')
				} else {
					PrintStr(Itoa(res))
				}
			} else if args[1] == "%" {
				res = num1 % num2
				PrintStr(Itoa(res))
			}
		} else {
			z01.PrintRune('0')
			z01.PrintRune(10)
		}
	}
}

// go run main.go  -9223372036854775808 "*" 0
