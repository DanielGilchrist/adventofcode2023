package day2

import (
	"bufio"
	"main/cmd/utils"
	"strconv"
	"strings"
)

const (
	Blue  = "blue"
	Red   = "red"
	Green = "green"
)

type Game struct {
	ID    int
	Blue  []int
	Red   []int
	Green []int
}

func (game Game) IsValid(maxRed, maxGreen, maxBlue int) bool {
	return game.belowMax(Blue, maxBlue) && game.belowMax(Red, maxRed) && game.belowMax(Green, maxGreen)
}

func (game Game) colourTurns(colour string) []int {
	switch colour {
	case Blue:
		return game.Blue
	case Red:
		return game.Red
	case Green:
		return game.Green
	default:
		return []int{}
	}
}

func (game Game) belowMax(colour string, max int) bool {
	turns := game.colourTurns(colour)

	for _, turn := range turns {
		if turn > max {
			return false
		}
	}

	return true
}

func NewGame(id int, turns [][2]string) Game {
	game := Game{
		ID: id,
	}

	for _, turn := range turns {
		num := stringToInt(turn[0])
		colour := turn[1]

		switch colour {
		case Blue:
			game.Blue = append(game.Blue, num)
		case Red:
			game.Red = append(game.Red, num)
		case Green:
			game.Green = append(game.Green, num)
		}
	}

	return game
}

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
