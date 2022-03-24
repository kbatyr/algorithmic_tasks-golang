package student

func Atoi(s string) int {
	var res int
	s1 := s
	if len(s) == 0 {
		return 0
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
				return 0
			}
		}
		if s1[0] == '-' {
			res = -res
		}
		if (s1[0] == '-' && res > 0) || (s1[0] == '+' || (s1[0] >= '0' && s1[0] <= '9')) && res < 0 {
			return 0
		}
	}
	return res
}
