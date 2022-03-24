package student

func RecursivePower(nb int, power int) int {
	var res int

	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}

	res = nb * RecursivePower(nb, power-1)
	return res
}
