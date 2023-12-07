package day3

import (
	"bufio"
	"fmt"
	"main/cmd/utils"
	"strconv"
	"unicode"
)

func Part1() {
	grid := parseInput()

	var sum int
	for outeri, row := range grid {
		var currentNumberStr []rune
		var validNumber bool

		for outerj, c := range row {
			if unicode.IsDigit(c) {
				currentNumberStr = append(currentNumberStr, c)
				indicesToCheck := [][2]int{
					{outeri, outerj - 1},
					{outeri, outerj + 1},
					{outeri - 1, outerj},
					{outeri - 1, outerj - 1},
					{outeri - 1, outerj + 1},
					{outeri + 1, outerj},
					{outeri + 1, outerj - 1},
					{outeri + 1, outerj + 1},
				}

				for _, pair := range indicesToCheck {
					i := pair[0]
					j := pair[1]

					if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) {
						if isSymbol(rune(grid[i][j])) {
							validNumber = true
						}
					}
				}
			}

			if len(row) == outerj+1 || !unicode.IsDigit(c) {
				if validNumber {
					num := stringToInt(string(currentNumberStr))
					sum += num
				}

				currentNumberStr = nil
				validNumber = false
			}
		}
	}

	fmt.Println(sum)
}

func parseInput() []string {
	var output []string
	utils.ReadInputFile("day3.txt", func(scanner *bufio.Scanner) {
		output = append(output, scanner.Text())
	})
	return output
}

func isSymbol(c rune) bool {
	return c != '.' && !unicode.IsDigit(c)
}

func isPeriod(c rune) bool {
	return c == '.'
}

func stringToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}
