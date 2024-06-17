package piscine

func TrimAtoi(s string) int {
	var num int
	negCheck := false

	for _, v := range s {
		if v == '-' && num == 0 {
			negCheck = true
			continue
		}

		switch v {
		case '0':
			num = num * 10
		case '1':
			num = (num + 1) * 10
		case '2':
			num = (num + 2) * 10
		case '3':
			num = (num + 3) * 10
		case '4':
			num = (num + 4) * 10
		case '5':
			num = (num + 5) * 10
		case '6':
			num = (num + 6) * 10
		case '7':
			num = (num + 7) * 10
		case '8':
			num = (num + 8) * 10
		case '9':
			num = (num + 9) * 10
		}
	}

	if num == 0 {
		return 0
	}

	num = num / 10
	if negCheck {
		num = -num
	}

	return num
}
