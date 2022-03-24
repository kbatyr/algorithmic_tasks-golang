package student

func Split(s, sep string) []string {
	var res []string
	if sep == "" {
		for _, ch := range s {
			res = append(res, string(ch))
		}
	} else {
		for i := 0; i <= len(s)-len(sep); i++ {

			if s[i:len(sep)+i] == sep {
				s = s[:i] + " " + s[len(sep)+i:]
			}
		}

		res = append(res, s)
	}

	return res
}
