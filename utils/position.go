package utils

//Position a position
type Position struct {
	X int
	Y int
}

//New Constructor
func New(x, y int) Position {
	return Position{x, y}
}
