package parser

import (
	"strconv"
	"strings"
)

const BlueCube string = "blue"
const RedCube string = "red"
const GreenCube string = "green"
const gameString string = "Game"

func parseRoundToInt(color string, round string) (int, error) {
	round = strings.ReplaceAll(round, color, "")
	round = strings.Trim(round, " ")
	return strconv.Atoi(round)
}

func ParseCubeToNumber(cubeNumber string, cubeColor string) (int, error) {
	// TODO: verify that the format of the cube number is correct
	cubeNumber = strings.Replace(cubeNumber, "=", " ", 1)
	return parseRoundToInt(cubeColor, cubeNumber)
}
