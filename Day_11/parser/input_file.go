package parser

import (
	"bufio"
	"io"
	"slices"
	"strings"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_11/galaxy"
)

const emptyItemsToAdd int = 1000000 - 1

func ParseToGalaxies(file io.Reader) galaxy.Galaxy {
	var galaxiesTmp galaxy.GalaxyMap
	var emptyRowsIndex, emptyColsIndex []int
	scanner := bufio.NewScanner(file)
	var i int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		row := make([]string, len(line))
		for i, c := range line {
			row[i] = string(c)
		}
		galaxiesTmp = append(galaxiesTmp, row)
		line = strings.ReplaceAll(line, ".", "")
		if len(line) == 0 {
			emptyRowsIndex = append(emptyRowsIndex, i)
		}
		i++
	}
	for j := 0; j < len(galaxiesTmp[0]); j++ {
		var countNonEmpty int
		for i := 0; i < len(galaxiesTmp); i++ {
			if galaxiesTmp[i][j] != "." {
				countNonEmpty++
			}
		}
		if countNonEmpty == 0 {
			emptyColsIndex = append(emptyColsIndex, j)
		}
	}
	var locations []galaxy.Point
	var addToRowIndex int
	for i, v := range galaxiesTmp {
		if slices.Contains(emptyRowsIndex, i) {
			addToRowIndex += emptyItemsToAdd
		}
		var addToIndex int
		for j := 0; j < len(v); j++ {
			if slices.Contains(emptyColsIndex, j) {
				addToIndex += emptyItemsToAdd
			}
			if v[j] == "#" {
				locations = append(locations, galaxy.Point{I: i + addToRowIndex, J: j + addToIndex})
			}
		}
	}
	return galaxy.Galaxy{
		GalaxyLocation: locations,
	}
}
