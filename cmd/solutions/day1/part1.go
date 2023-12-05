package day1

func Part1() {
	calibrate(func(line string) string {
		return extractDigits(line)
	})
}
