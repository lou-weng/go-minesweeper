package enums

type Difficulty int

const (
	Easy	Difficulty = (iota + 1) * 15 
	Intermediate
	Hard
)

func (difficulty Difficulty) String() string {
	strings := [...]string{
		"Easy",
		"Intermediate",
		"Hard",
	}

	if difficulty < Easy || difficulty > Hard {
		return "Unknown"
	}

	return strings[difficulty]
}
