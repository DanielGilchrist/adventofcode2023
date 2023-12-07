package day2

import "fmt"

func Part2() {
	games := parseGames()
	colourMaps := gamesToColourMaps(games)

	var sum int
	for _, colourMap := range colourMaps {
		power := 1

		for _, colourNum := range colourMap {
			power *= colourNum
		}

		sum += power
	}

	fmt.Println(sum)
}

func gamesToColourMaps(games []Game) []map[string]int {
	var colourMaps []map[string]int

	for _, game := range games {
		colourMaps = append(colourMaps, game.ToLowestPossibleColoursMap())
	}

	return colourMaps
}
