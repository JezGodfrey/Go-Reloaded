package piscine

func ToUpper(s string) string {
	var upString []byte

	for _, v := range []byte(s) {
		if v > 96 && v < 123 {
			upString = append(upString, v-32)
		} else {
			upString = append(upString, v)
		}
	}

	return string(upString)
}
