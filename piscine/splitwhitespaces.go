package piscine

func SplitWhiteSpaces(s string) []string {
	var ss []string
	var newString []rune

	for _, v := range s {
		if v == 9 || v == 10 || v == 32 {
			if string(newString) != "" {
				ss = append(ss, string(newString))
			}
			newString = nil
		} else {
			newString = append(newString, v)
		}
	}

	if newString != nil {
		ss = append(ss, string(newString))
	}

	return ss
}
