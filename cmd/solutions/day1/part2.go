package day1

import (
	"fmt"
	"strings"
)

var wordsToDigits = map[string]string{
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
	calibrate(func(line string) string {
		converted := convertWordsToDigits(line)
		return extractDigits(converted)
	})
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
