package day2

import (
	"bufio"
	"main/cmd/utils"
	"strconv"
	"strings"
)

func parseGames() []Game {
	var games []Game
	gameId := 1

	utils.ReadInputFile("day2.txt", func(scanner *bufio.Scanner) {
		// "Game 1: 1 red, 5 blue, 10 green; 5 green, 6 blue, 12 red; 4 red, 10 blue, 4 green"
		line := scanner.Text()

		// ["Game 1", "1 red, 5 blue, 10 green; 5 green, 6 blue, 12 red; 4 red, 10 blue, 4 green"]
		split := strings.Split(line, ": ")

		// "1 red, 5 blue, 10 green, 5 green, 6 blue, 12 red, 4 red, 10 blue, 4 green"
		rest := strings.ReplaceAll(split[1], "; ", ", ")

		var turns [][2]string

		// ["1 red", "5 blue", "10 green", "5 green", "6 blue", "12 red", "4 red", "10 blue", "4 green"]
		for _, gameStr := range strings.Split(rest, ", ") {
			// ["1", "red"]
			gameStrSplit := strings.Split(gameStr, " ")

			// "1"
			numColour := gameStrSplit[0]

			// "red"
			colour := gameStrSplit[1]

			turns = append(turns, [2]string{numColour, colour})
		}

		games = append(games, NewGame(gameId, turns))
		gameId++
	})

	return games
}

func stringToInt(digit string) int {
	num, err := strconv.Atoi(digit)

	if err != nil {
		return -1
	}

	return num
}
