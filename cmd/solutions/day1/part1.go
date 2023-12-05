package day1

import (
	"bufio"
	"fmt"
	"strconv"

	"main/cmd/utils"
)

func Part1() {
	file := utils.OpenInputFile("day1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int

	for scanner.Scan() {
		digits := extractDigits(scanner.Text())

		if len(digits) > 0 {
			firstDigit := string(digits[0])
			lastDigit := string(digits[len(digits)-1])
			num, err := strconv.Atoi(firstDigit + lastDigit)

			if err == nil {
				sum += num
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sum)
}
