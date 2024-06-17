package piscine

func Capitalize(s string) string {
	var cap []byte
	firstLetter := true

	for _, v := range []byte(s) {
		if (v < 48 || v > 57) && (v < 65 || v > 90) && (v < 97 || v > 122) {
			firstLetter = true
		}

		if firstLetter {
			if (v > 64 && v < 91) || (v > 47 && v < 58) {
				firstLetter = false
			}
			if v > 96 && v < 123 {
				cap = append(cap, v-32)
				firstLetter = false
			} else {
				cap = append(cap, v)
			}
		} else {
			if v > 64 && v < 91 {
				cap = append(cap, v+32)
			} else {
				cap = append(cap, v)
			}
		}
	}

	return string(cap)
}
