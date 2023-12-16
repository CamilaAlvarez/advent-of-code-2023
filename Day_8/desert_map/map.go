package desertmap

import (
	"log"
	"strings"

	mymath "github.com/CamilaAlvarez/advent-of-code-2023/Day_8/math"
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

// Paths reach a Z ending path in a periodic manner. Finding the minimum number of steps it takes
// each path, and finding the LCM gives us the total number of steps
func (m MapFile) NumberOfSteps(from []Node) int {
	var count, completed int
	var zSuffix string = "Z"
	currentNodes := from
	stepsForNode := make([]int, len(from))
	for m.Pattern.HasNext() {
		p := m.Pattern.GetNext()
		count++
		for i, n := range currentNodes {
			v, ok := m.Map[n]
			if !ok {
				log.Fatal("Invalid node: ", n)
			}
			switch p {
			case L:
				currentNodes[i] = v.L
			case R:
				currentNodes[i] = v.R
			default:
				log.Fatal("Invalid pattern symbol:", p)
			}
			if strings.HasSuffix(string(currentNodes[i]), zSuffix) && stepsForNode[i] == 0 {
				stepsForNode[i] = count
				completed++
			}
		}
		if completed == len(currentNodes) {
			break
		}
	}
	lcm := (stepsForNode[0] * stepsForNode[1]) / mymath.GCD(stepsForNode[0], stepsForNode[1])
	for i := 2; i < len(stepsForNode); i++ {
		lcm = (lcm * stepsForNode[i]) / mymath.GCD(lcm, stepsForNode[i])
	}
	return lcm
}
func (m MapFile) StartingPoints() []Node {
	aSuffix := "A"
	var startingPoints []Node
	for k := range m.Map {
		if strings.HasSuffix(string(k), aSuffix) {
			startingPoints = append(startingPoints, k)
		}
	}
	return startingPoints
}
