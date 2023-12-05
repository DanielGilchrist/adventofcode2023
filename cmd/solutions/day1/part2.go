package day1

import (
	"bufio"
	"fmt"
	"main/cmd/utils"
	"strconv"
	"strings"
)

var wordsToDigits = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func Part2() {
	file := utils.OpenInputFile("day1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int

	for scanner.Scan() {
		line := convertWordsToDigits(scanner.Text())
		digits := extractDigits(line)

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

func convertWordsToDigits(line string) string {
	for word, number := range wordsToDigits {
		first := string(word[0])
		last := string(word[len(word)-1])
		paddedNumber := fmt.Sprintf("%s%s%s", first, number, last)
		line = strings.ReplaceAll(line, word, paddedNumber)
	}

	return line
}
