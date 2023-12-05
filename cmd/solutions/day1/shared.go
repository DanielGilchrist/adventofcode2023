package day1

func extractDigits(line string) string {
	digits := ""
	for _, c := range line {
		if c >= '0' && c <= '9' {
			digits += string(c)
		}
	}
	return digits
}
