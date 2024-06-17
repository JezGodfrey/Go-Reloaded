package piscine

func DeleteElement(ss []string, i int) []string {
	for j := i; j < len(ss)-1; j++ {
		ss[j] = ss[j+1]
	}
	ss = ss[:len(ss)-1]

	return ss
}
