package desertmap

import (
	"log"
	"math"
)

type Direction rune
type Node string

const L = 'L'
const R = 'R'

type NodeMapping struct {
	L Node
	R Node
}
type MapFile struct {
	Pattern *InstructionPattern
	Map     map[Node]NodeMapping
}

type InstructionPattern struct {
	Pattern string
	index   int
}

func (ip *InstructionPattern) HasNext() bool {
	return true
}
func (ip *InstructionPattern) GetNext() Direction {
	nextDirection := Direction(ip.Pattern[ip.index])
	ip.index = (ip.index + 1) % len(ip.Pattern)
	return nextDirection
}

func (m MapFile) NumberOfSteps(from Node, to Node) int {
	var count int
	if from == to {
		return count
	}
	currentNode := from
	for m.Pattern.HasNext() {
		p := m.Pattern.GetNext()
		v, ok := m.Map[currentNode]
		if !ok {
			log.Fatal("Invalid node: ", currentNode)
		}
		count++
		switch p {
		case L:
			currentNode = v.L
		case R:
			currentNode = v.R
		default:
			log.Fatal("Invalid pattern symbol:", p)
		}
		if currentNode == to {
			return count
		}
	}
	log.Fatal("No path between", from, "and to", to, "was found")
	return math.MaxInt
}
