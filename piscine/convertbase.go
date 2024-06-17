package piscine

func GetBaseNumbers(s string, bf string) []int {
	var numbers []int
	b := []byte(s)

	for _, v := range b {
		for i := 0; i < len(bf); i++ {
			if v == bf[i] {
				numbers = append(numbers, i)
			}
		}
	}

	return numbers
}

func RecursivePow(nb int, power int) int {
	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}

	if power == 1 {
		return nb
	}

	if power > 1 {
		return nb * RecursivePow(nb, power-1)
	}

	return 0
}

func BaseToNumber(b int, ns []int) int {
	var convNum int
	for i := 0; i < len(ns); i++ {
		convNum = convNum + (RecursivePow(b, i) * ns[i])
	}

	return convNum
}

func NumberToBase(n int, b string, r *[]rune) int {
	*r = append(*r, rune(b[n%len(b)]))
	if n%len(b) != n {
		return NumberToBase(n/len(b), b, r)
	}
	return 0
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	base := len(baseFrom)
	var revnbr string
	var runes []rune
	var revrunes []rune

	for i := len(nbr) - 1; i >= 0; i-- {
		revnbr = revnbr + string(nbr[i])
	}

	numbers := GetBaseNumbers(revnbr, baseFrom)
	cNum := BaseToNumber(base, numbers)
	NumberToBase(cNum, baseTo, &runes)

	for i := len(runes) - 1; i >= 0; i-- {
		revrunes = append(revrunes, runes[i])
	}

	return string(revrunes)
}
