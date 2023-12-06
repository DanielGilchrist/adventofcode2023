package cmd

import (
	"github.com/spf13/cobra"

	"main/cmd/solutions/day2"
)

var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "Day 2 solutions",
	Long:  "My day 2 solutions for Advent of code 2023",
	Run: func(cmd *cobra.Command, args []string) {
		if part1 {
			day2.Part1()
		} else if part2 {
			// day2.Part2()
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(day2Cmd)
}
