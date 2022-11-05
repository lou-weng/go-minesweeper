package enums

type GameTile int

const (
	Bomb	GameTile = iota + 1
	Blank
)

func (gameTile GameTile) String() string {
	strings := [...]string{
		"Bomb",
		"Blank",
	}

	if gameTile < Bomb || gameTile > Blank {
		return "Unknown"
	}

	return strings[gameTile]
}
