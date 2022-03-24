package student

// ActiveBits returns the number of active bits (bits with the value 1) in the binary representation of an integer number.
func ActiveBits(n int) uint {
	var mod, count uint
	for n != 0 {
		mod = uint(n % 2)
		if mod == 1 {
			count++
		}
		n = n / 2
	}

	return count
}
