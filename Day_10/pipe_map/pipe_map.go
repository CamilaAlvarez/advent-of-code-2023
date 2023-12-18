package pipemap

import (
	"fmt"
	"log"
)

type Pipe int
type OriginDirection int
type Location struct {
	I      int
	J      int
	Origin OriginDirection
}

const (
	_ OriginDirection = iota
	Up
	Down
	Left
	Right
)
const (
	S Pipe = iota
	I
	H
	L
	J
	Seven
	F
	Dot
)

type PipeMap struct {
	Map           [][]Pipe
	StartLocation Location
}

func (pm PipeMap) FindMaxDistance() int {
	// We find the two starting points from S, there may be more, so we will also include flags to indicate which paths to follow
	var locations []Location
	var pathIsBroken [4]bool
	if pm.StartLocation.I > 0 && pm.Map[pm.StartLocation.I-1][pm.StartLocation.J] != Dot {
		locations = append(locations, Location{I: pm.StartLocation.I - 1, J: pm.StartLocation.J, Origin: Down})
	}
	numberRows := len(pm.Map)
	if pm.StartLocation.I < numberRows-1 && pm.Map[pm.StartLocation.I+1][pm.StartLocation.J] != Dot {
		locations = append(locations, Location{I: pm.StartLocation.I + 1, J: pm.StartLocation.J, Origin: Up})
	}

	if pm.StartLocation.J > 0 && pm.Map[pm.StartLocation.J][pm.StartLocation.J-1] != Dot {
		locations = append(locations, Location{I: pm.StartLocation.I, J: pm.StartLocation.J - 1, Origin: Right})
	}

	if pm.StartLocation.J < numberRows-1 && pm.Map[pm.StartLocation.I][pm.StartLocation.J+1] != Dot {
		locations = append(locations, Location{I: pm.StartLocation.I, J: pm.StartLocation.J + 1, Origin: Left})
	}
	numSteps := 1
	for {
		for i, l := range locations {
			if pathIsBroken[i] {
				continue
			}
			pipeType := pm.Map[l.I][l.J]
			origin := l.Origin
			switch pipeType {
			case I:
				if origin == Down {
					locations[i] = Location{I: l.I - 1, J: l.J, Origin: Down}
				} else if origin == Up {
					locations[i] = Location{I: l.I + 1, J: l.J, Origin: Up}
				} else {
					fmt.Println("Invalid origin for |: ", origin)
					pathIsBroken[i] = true
				}
			case H:
				if origin == Left {
					locations[i] = Location{I: l.I, J: l.J + 1, Origin: Left}
				} else if origin == Right {
					locations[i] = Location{I: l.I, J: l.J - 1, Origin: Right}
				} else {
					fmt.Println("Invalid origin for -: ", origin)
					pathIsBroken[i] = true
				}
			case L:
				if origin == Up {
					locations[i] = Location{I: l.I, J: l.J + 1, Origin: Left}
				} else if origin == Right {
					locations[i] = Location{I: l.I - 1, J: l.J, Origin: Down}
				} else {
					fmt.Println("Invalid origin for L: ", origin)
					pathIsBroken[i] = true
				}
			case J:
				if origin == Up {
					locations[i] = Location{I: l.I, J: l.J - 1, Origin: Right}
				} else if origin == Left {
					locations[i] = Location{I: l.I - 1, J: l.J, Origin: Down}
				} else {
					fmt.Println("Invalid origin for J: ", origin)
					pathIsBroken[i] = true
				}
			case Seven:
				if origin == Down {
					locations[i] = Location{I: l.I, J: l.J - 1, Origin: Right}
				} else if origin == Left {
					locations[i] = Location{I: l.I + 1, J: l.J, Origin: Up}
				} else {
					fmt.Println("Invalid origin for 7: ", origin)
					pathIsBroken[i] = true
				}
			case F:
				if origin == Down {
					locations[i] = Location{I: l.I, J: l.J + 1, Origin: Left}
				} else if origin == Right {
					locations[i] = Location{I: l.I + 1, J: l.J, Origin: Up}
				} else {
					fmt.Println("Invalid origin for F: ", origin)
					pathIsBroken[i] = true
				}
			default:
				log.Fatal("Following an invalid path: ", pipeType)
			}
		}
		numSteps++
		if reachedEnd(locations, pathIsBroken) {
			break
		}
	}
	return numSteps
}

func reachedEnd(locations []Location, pathIsBroken [4]bool) bool {
	var location Location
	isFirstValidValue := true
	for i, v := range locations {
		if pathIsBroken[i] {
			continue
		}
		if isFirstValidValue {
			location = v
			isFirstValidValue = false
		} else if location.I != v.I || location.J != v.J {
			return false
		}
	}
	return true
}
