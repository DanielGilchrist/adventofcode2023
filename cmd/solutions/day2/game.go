package day2

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

func (game Game) IsValid(maxRed, maxGreen, maxBlue int) bool {
	return game.belowMax(Blue, maxBlue) && game.belowMax(Red, maxRed) && game.belowMax(Green, maxGreen)
}

func (game Game) ToLowestPossibleColoursMap() map[string]int {
	colours := []string{Blue, Red, Green}
	colourMap := map[string]int{}

	for _, colour := range colours {
		colourMap[colour] = game.toLowestPossibleColour(colour)
	}

	return colourMap
}

func (game Game) toLowestPossibleColour(colour string) int {
	turns := game.colourTurns(colour)
	highest := turns[0]

	for _, turn := range turns {
		if turn > highest {
			highest = turn
		}
	}

	return highest
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
