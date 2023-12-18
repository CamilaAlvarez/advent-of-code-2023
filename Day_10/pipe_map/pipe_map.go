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
type Cell int

const (
	Outside Cell = iota
	Inside
	Border
	Start
	ClosedBorder
)
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
	typesCells := make([][]Cell, len(pm.Map))
	for i, v := range pm.Map {
		typesCells[i] = make([]Cell, len(v))
	}
	typesCells[pm.StartLocation.I][pm.StartLocation.J] = Start
	if pm.StartLocation.I > 0 && pm.Map[pm.StartLocation.I-1][pm.StartLocation.J] != Dot {
		locations = append(locations, Location{I: pm.StartLocation.I - 1, J: pm.StartLocation.J, Origin: Down})
	}
	numberRows := len(pm.Map)
	if pm.StartLocation.I < numberRows-1 && pm.Map[pm.StartLocation.I+1][pm.StartLocation.J] != Dot {
		locations = append(locations, Location{I: pm.StartLocation.I + 1, J: pm.StartLocation.J, Origin: Up})
	}
	if pm.StartLocation.J > 0 && pm.Map[pm.StartLocation.I][pm.StartLocation.J-1] != Dot {
		locations = append(locations, Location{I: pm.StartLocation.I, J: pm.StartLocation.J - 1, Origin: Right})
	}

	if pm.StartLocation.J < numberRows-1 && pm.Map[pm.StartLocation.I][pm.StartLocation.J+1] != Dot {
		locations = append(locations, Location{I: pm.StartLocation.I, J: pm.StartLocation.J + 1, Origin: Left})
	}

	numSteps := 1
	for {
		for i, l := range locations {
			oldLocation := l
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
					continue
				}
			case H:
				if origin == Left {
					locations[i] = Location{I: l.I, J: l.J + 1, Origin: Left}
				} else if origin == Right {
					locations[i] = Location{I: l.I, J: l.J - 1, Origin: Right}
				} else {
					fmt.Println("Invalid origin for -: ", origin)
					pathIsBroken[i] = true
					continue
				}
			case L:
				if origin == Up {
					locations[i] = Location{I: l.I, J: l.J + 1, Origin: Left}
				} else if origin == Right {
					locations[i] = Location{I: l.I - 1, J: l.J, Origin: Down}
				} else {
					fmt.Println("Invalid origin for L: ", origin)
					pathIsBroken[i] = true
					continue
				}
			case J:
				if origin == Up {
					locations[i] = Location{I: l.I, J: l.J - 1, Origin: Right}
				} else if origin == Left {
					locations[i] = Location{I: l.I - 1, J: l.J, Origin: Down}
				} else {
					fmt.Println("Invalid origin for J: ", origin)
					pathIsBroken[i] = true
					continue
				}
			case Seven:
				if origin == Down {
					locations[i] = Location{I: l.I, J: l.J - 1, Origin: Right}
				} else if origin == Left {
					locations[i] = Location{I: l.I + 1, J: l.J, Origin: Up}
				} else {
					fmt.Println("Invalid origin for 7: ", origin)
					pathIsBroken[i] = true
					continue
				}
			case F:
				if origin == Down {
					locations[i] = Location{I: l.I, J: l.J + 1, Origin: Left}
				} else if origin == Right {
					locations[i] = Location{I: l.I + 1, J: l.J, Origin: Up}
				} else {
					fmt.Println("Invalid origin for F: ", origin)
					pathIsBroken[i] = true
					continue
				}
			default:
				log.Fatal("Following an invalid path: ", pipeType)
			}
			typesCells[oldLocation.I][oldLocation.J] = Border
		}
		numSteps++
		if reachedEnd(locations, pathIsBroken) {
			for i, v := range locations {
				if pathIsBroken[i] {
					continue
				}
				typesCells[v.I][v.J] = Border
				break
			}
			break
		}
	}
	var elementsInsideCycle int
	fmt.Println(typesCells)
	for i := 0; i < len(typesCells); i++ {
		var j int
		seenBorder := false
		for j < len(typesCells[i]) {
			//fmt.Printf("i=%d, j=%d\n", i, j)
			//fmt.Println(pm.Map[i][j])

			if typesCells[i][j] == Border {
				pipeType := pm.Map[i][j]
				seenBorder = true
			jLoop:
				for {
					j++
					if j >= len(typesCells[i]) {
						break
					}
					if i > 0 && i < len(typesCells)-1 {
						switch pipeType {
						case S:
							continue
						case I:
							if typesCells[i][j] == Border && (pm.Map[i][j] == I || pm.Map[i][j] == F || pm.Map[i][j] == L) {
								//fmt.Printf("i=%d, j=%d\n", i, j)
								//fmt.Println(pm.Map[i][j])
								//typesCells[i][j] = ClosedBorder
								seenBorder = false
								j++
								break jLoop
							} else if typesCells[i][j] != Border {
								typesCells[i][j] = Inside
								//fmt.Printf("i=%d, j=%d\n", i, j)
								//fmt.Println(pm.Map[i][j])
								elementsInsideCycle++
							}
						case H:
							continue
						case L:
							if typesCells[i][j] == Border && (pm.Map[i][j] == J || pm.Map[i][j] == Seven) {
								//typesCells[i][j] = ClosedBorder
								j++
								seenBorder = false
								break jLoop
							} else if typesCells[i][j] != Border {
								typesCells[i][j] = Inside
								//fmt.Printf("i=%d, j=%d\n", i, j)
								elementsInsideCycle++
							}
						case J:
							if typesCells[i][j] == Border && (pm.Map[i][j] == F || pm.Map[i][j] == I || pm.Map[i][j] == L) {
								//typesCells[i][j] = ClosedBorder
								j++
								seenBorder = false
								break jLoop
							} else if typesCells[i][j] != Border {
								typesCells[i][j] = Inside
								//fmt.Printf("i=%d, j=%d\n", i, j)
								elementsInsideCycle++
							}
						case Seven:
							if typesCells[i][j] == Border && (pm.Map[i][j] == F || pm.Map[i][j] == I || pm.Map[i][j] == L) {
								//typesCells[i][j] = ClosedBorder
								j++
								seenBorder = false
								break jLoop
							} else if typesCells[i][j] != Border {
								typesCells[i][j] = Inside
								fmt.Printf("i=%d, j=%d\n", i, j)
								elementsInsideCycle++
							}
						case F:
							if typesCells[i][j] == Border && (pm.Map[i][j] == Seven || pm.Map[i][j] == J) {
								//typesCells[i][j] = ClosedBorder
								j++
								seenBorder = false
								break jLoop
							} else if typesCells[i][j] != Border {
								typesCells[i][j] = Inside
								//fmt.Printf("i=%d, j=%d\n", i, j)
								elementsInsideCycle++
							}

						}
					}
				}
			} else if seenBorder && i > 0 && j > 0 && i < len(typesCells)-1 && j < len(typesCells[i])-1 && typesCells[i][j] != Border && (typesCells[i-1][j] == Border || typesCells[i-1][j] == Inside) && (typesCells[i][j-1] == Border || typesCells[i][j-1] == Inside) {
				startIndex := j
				endIndex := j
				var reachedBorder bool
				for j < len(typesCells[i])-1 && typesCells[i][j] != Border {
					endIndex++
					j++
					if typesCells[i][j] == Border {
						reachedBorder = true
					}
				}
				if reachedBorder {
					for k := startIndex; k < endIndex; k++ {
						typesCells[i][k] = Inside
						elementsInsideCycle++
					}
				}
			} else {
				j++
			}
		}
	}
	fmt.Println(elementsInsideCycle)
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
