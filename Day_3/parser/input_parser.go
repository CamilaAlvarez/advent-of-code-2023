package parser

import (
	"bufio"
	"io"
)

func ParseInputToMatrix(file io.Reader) ([][]rune, [][]bool) {
	var matrix [][]rune
	var collisionMatrix [][]bool
	scanner := bufio.NewScanner(file)
	// TODO: check scanner error
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 0 {
			continue
		}
		row := make([]rune, 0, len(line))
		for _, v := range line {
			row = append(row, v)
		}
		matrix = append(matrix, row)
		collisionMatrix = append(collisionMatrix, make([]bool, len(row)))
	}
	return matrix, collisionMatrix
}
