package student

import (
	"github.com/01-edu/z01"
)

func printNbrBase(nbr int, base string) string {
	var count1, count2, mod int
	var rev, res string
	v := 1

	for _, let1 := range base {
		if let1 == '+' || let1 == '-' {
			count1++
		}
		for _, let2 := range base {
			if let1 == let2 {
				count2++
			}
		}
	}
	if len(base) < 2 || count1 > 0 || count2 > len(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else {
		if nbr < 0 {
			z01.PrintRune('-')
			v = -1
		}

		for nbr != 0 {
			mod = nbr % len(base) * v
			rev += string(base[mod])
			nbr = nbr / len(base)
		}

		for i := len(rev) - 1; i >= 0; i-- {
			res += string(rev[i])
		}
	}
	return res
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	dec := AtoiBase(nbr, baseFrom)   // dec int
	res := printNbrBase(dec, baseTo) // res string
	return res
}
