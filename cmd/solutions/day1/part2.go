package day1

import (
	"bufio"
	"fmt"
	"main/cmd/utils"
	"slices"
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

	var i int
	for scanner.Scan() {
		fmt.Println("i: ", i)
		text := scanner.Text()
		fmt.Println("1: ", text)
		line := convertWordsToDigits(text)
		digits := extractDigits(line)
		fmt.Println("2: ", digits)

		if len(digits) > 0 {
			firstDigit := string(digits[0])
			lastDigit := string(digits[len(digits)-1])
			num, err := strconv.Atoi(firstDigit + lastDigit)

			if err == nil {
				fmt.Println("3: ", num)
				sum += num
			}
		}

		i++
		fmt.Println("")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	i++
	fmt.Println(sum)
}

func convertWordsToDigits(line string) string {
	var indexMap = make(map[int]string)

	for word, number := range wordsToDigits {
		if index := strings.Index(line, word); index != -1 {
			indexMap[index] = number
		}
	}

	i := 0
	keys := sortedKeys(indexMap)
	for _, index := range keys {
		number := indexMap[index]
		offsetIndex := index + i
		line = line[:offsetIndex] + number + line[offsetIndex:]
		i++
	}

	return line
}

func sortedKeys(m map[int]string) []int {
	keys := make([]int, len(m))
	i := 0

	for k := range m {
		keys[i] = k
		i++
	}

	slices.Sort(keys)

	return keys
}
