package student

func SplitWhiteSpaces(s string) []string {
	var res []string
	var word string
	for _, let := range s {

		if let != ' ' && let != '	' && let != '\n' {
			word += string(let)
		} else {
			if word != "" {
				res = append(res, word)
				word = ""
			}
		}
	}
	if word != "" {
		res = append(res, word)
	}
	return res
}
