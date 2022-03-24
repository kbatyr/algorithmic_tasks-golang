package student

func AtoiBase(s string, base string) int {
	var count, res int
	var rev string
	for _, let1 := range base {
		if let1 == '+' || let1 == '-' {
			return 0
		}
		for _, let2 := range base {
			if let1 == let2 {
				count++
			}
		}
	}

	if len(base) < 2 || count > len(base) {
		return 0
	} else {
		for i := len(s) - 1; i >= 0; i-- {
			rev += string(s[i])
		}
		for i1, let1 := range rev {
			for i2, let2 := range base {
				if let1 == let2 {
					res += i2 * RecursivePower(len(base), i1)
				}
			}
		}
	}
	return res
}
