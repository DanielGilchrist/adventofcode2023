/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/spf13/cobra"
)

var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Day 1 solutions",
	Long:  "My day 1 solutions for Advent of code 2023",
	Run: func(cmd *cobra.Command, args []string) {
		if part1 {
			Part1()
		} else if part2 {
			Part2()
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)
}

func OpenFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return file
}

func Part1() {
	file := OpenFile("./cmd/inputs/day1.txt")
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

func Part2() {
	fmt.Println("TODO")
}
