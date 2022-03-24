package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) >= 2 {
		args := os.Args[1:]
		var arr1 []rune
		var arr2 []int
		vowels := "AEIOUaeiou"

		// Заносим все аргументы с консоли в один срез, разделяя их пробелами
		for i1, elems := range args {
			for _, let := range elems {
				arr1 = append(arr1, let)
			}
			if i1 != len(args)-1 {
				arr1 = append(arr1, ' ')
			}
		}
		// Заносим в срез2 индексы глассных букв в срезе1
		for i2, let := range arr1 {
			for _, vow := range vowels {
				if let == vow {
					arr2 = append(arr2, i2)
				}
			}
		}

		// цикл работающий до середины среза2
		// меняем местами глассные буквы в срезе1 с пом. ранее полученных индексов глассных в срез2
		for i := 0; i < len(arr2)/2; i++ {
			arr1[arr2[i]], arr1[arr2[len(arr2)-1-i]] = arr1[arr2[len(arr2)-1-i]], arr1[arr2[i]]
		}

		// печатаем изменённый срез1
		for _, let := range arr1 {
			z01.PrintRune(let)
		}
		z01.PrintRune(10)
	} else {
		z01.PrintRune(10)
	}
}
