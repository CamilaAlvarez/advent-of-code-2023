package parser

import (
	"bufio"
	"io"
	"slices"
	"strings"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_11/galaxy"
)

func ParseToGalaxies(file io.Reader) galaxy.Galaxy {
	var galaxiesTmp, galaxies galaxy.GalaxyMap
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
	for i, v := range galaxiesTmp {
		if slices.Contains(emptyRowsIndex, i) {
			newRow := make([]string, len(v)+len(emptyColsIndex))
			for j := 0; j < len(v)+len(emptyColsIndex); j++ {
				newRow[j] = "."
			}
			galaxies = append(galaxies, newRow)
		}
		newRow := make([]string, len(v)+len(emptyColsIndex))
		var addToIndex int
		for j := 0; j < len(v); j++ {
			if slices.Contains(emptyColsIndex, j) {
				newRow[j+addToIndex] = "."
				addToIndex++
			}
			newRow[j+addToIndex] = v[j]
			if newRow[j+addToIndex] == "#" {
				locations = append(locations, galaxy.Point{I: len(galaxies), J: j + addToIndex})
			}
		}
		galaxies = append(galaxies, newRow)
	}
	return galaxy.Galaxy{
		Map:            galaxies,
		GalaxyLocation: locations,
	}
}
