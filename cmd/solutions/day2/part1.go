package day2

import (
	"fmt"
)

func Part1() {
	games := parseGames()

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	var validGames []Game

	for _, game := range games {
		if game.IsValid(maxRed, maxGreen, maxBlue) {
			validGames = append(validGames, game)
		}
	}

	var idSum int

	for _, game := range validGames {
		idSum += game.ID
	}

	fmt.Println(idSum)
}
