package student

// func Compare(a, b string) int {
// 	if a == b {
// 		return 0
// 	}
// 	if a < b {
// 		return -1
// 	}
// 	return +1
// }

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for range a {
		for i := range a {
			if i != len(a)-1 {
				if f(a[i], a[i+1]) == +1 {
					a[i], a[i+1] = a[i+1], a[i]
				}
			}
		}
	}
}
