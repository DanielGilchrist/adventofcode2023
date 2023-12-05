/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var part1 bool
var part2 bool

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "Advent of code 2023 CLI",
	Long:  "CLI app with my solutions for Advent of code 2023",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&part1, "part1", "1", false, "Execute part 1 solution")
	rootCmd.PersistentFlags().BoolVarP(&part2, "part2", "2", false, "Execute part 2 solution")
}
