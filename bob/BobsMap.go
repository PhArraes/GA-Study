package bob

import (
	"math"

	utils "../utils"
)

const north = 0
const south = 1
const east = 2
const west = 3

//BobsMap is a definition of a map
type BobsMap struct {
	startX int
	startY int
	endX   int
	endY   int
	bobMap [][]int
}

//New Constructs BobsMap
func New() BobsMap {
	bm := BobsMap{14, 7, 0, 2, bobMap()}
	return bm
}

//TestRoute Test a route and return its fitness
func (bobsMap BobsMap) TestRoute(steps []int) float64 {
	count := len(steps)
	currentPos := utils.New(bobsMap.startX, bobsMap.startY)
	for index := 0; index < count; index++ {
		currentPos = move(bobsMap.bobMap, currentPos, steps[index])
	}
	return bobsMap.calcFitness(currentPos)
}

func (bobsMap BobsMap) calcFitness(position utils.Position) float64 {
	diffX := math.Abs((float64)(position.X - bobsMap.endX))
	diffY := math.Abs((float64)(position.Y - bobsMap.endY))
	return 1 / (diffX + diffY + 1)
}

func move(bobsMap [][]int, currentPos utils.Position, move int) utils.Position {
	newPosition := newPosition(currentPos, move)
	if validPosition(newPosition, bobsMap) {
		return newPosition
	}
	return utils.New(currentPos.X, currentPos.Y)
}

func newPosition(position utils.Position, move int) utils.Position {
	switch move {
	case north:
		return utils.New(position.X, position.Y-1)
	case south:
		return utils.New(position.X, position.Y+1)
	case east:
		return utils.New(position.X+1, position.Y)
	case west:
		return utils.New(position.X-1, position.Y)
	}
	return position
}

func validPosition(position utils.Position, bobsMap [][]int) bool {
	field := bobsMap[position.X][position.Y]
	return field == 0 || field == 8
}

func bobMap() [][]int {
	return [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1},
		{8, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 5},
		{1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
}
