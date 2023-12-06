package cmd

import (
	"github.com/spf13/cobra"

	"main/cmd/solutions/day1"
)

var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Day 1 solutions",
	Long:  "My day 1 solutions for Advent of code 2023",
	Run: func(cmd *cobra.Command, args []string) {
		if part1 {
			day1.Part1()
		} else if part2 {
			day1.Part2()
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)
}
