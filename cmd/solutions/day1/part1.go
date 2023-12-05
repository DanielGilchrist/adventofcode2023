package day1

import (
	"bufio"
	"fmt"
	"strconv"
	"unicode"

	"main/cmd/utils"
)

func Part1() {
	file := utils.OpenInputFile("day1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int

	for scanner.Scan() {
		var firstAndLast [2]byte
		var lastChar byte

		for _, c := range scanner.Bytes() {
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
	}

	fmt.Println(sum)
}
