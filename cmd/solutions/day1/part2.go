package day1

import (
	"bufio"
	"fmt"
	"main/cmd/utils"
	"strconv"
	"strings"
	"unicode"
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
		var firstAndLast [2]byte
		var lastChar byte

		line := []byte(convertWordsToDigits(scanner.Text()))

		for _, c := range line {
			if unicode.IsDigit(rune(c)) {
				lastChar = c
			}

			if lastChar != 0 && firstAndLast[0] == 0 {
				firstAndLast[0] = lastChar
			}
		}

		firstAndLast[1] = lastChar

		num, e := strconv.Atoi(fmt.Sprintf("%s", firstAndLast[:]))

		if e == nil {
			sum += num
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sum)
}

// TODO: This doesn't handle cases where word numbers overlap e.g. twone or threeight
func convertWordsToDigits(line string) string {
	for word, number := range wordsToDigits {
		line = strings.Replace(line, word, number, -1)
	}

	return line
}
