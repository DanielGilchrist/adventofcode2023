package day1

import (
	"bufio"
	"fmt"
	"main/cmd/utils"
	"strconv"
)

func extractDigits(line string) string {
	digits := ""
	for _, c := range line {
		if c >= '0' && c <= '9' {
			digits += string(c)
		}
	}
	return digits
}

func calibrate(f func(string) string) {
	file := utils.OpenInputFile("day1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int

	for scanner.Scan() {
		digits := f(scanner.Text())

		if len(digits) > 0 {
			firstDigit := string(digits[0])
			lastDigit := string(digits[len(digits)-1])
			num, err := strconv.Atoi(firstDigit + lastDigit)

			if err == nil {
				sum += num
			}
		}
	}

	fmt.Println(sum)
}
