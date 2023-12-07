package cmd

import (
	"github.com/spf13/cobra"

	"main/cmd/solutions/day3"
)

var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "Day 3 solutions",
	Long:  "My day 3 solutions for Advent of code 2023",
	Run: func(cmd *cobra.Command, args []string) {
		if part1 {
			day3.Part1()
		} else if part2 {
			// day3.Part2()
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(day3Cmd)
}
