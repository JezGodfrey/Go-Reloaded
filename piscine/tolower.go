package piscine

func ToLower(s string) string {
	var lowString []byte

	for _, v := range []byte(s) {
		if v > 64 && v < 91 {
			lowString = append(lowString, v+32)
		} else {
			lowString = append(lowString, v)
		}
	}

	return string(lowString)
}
