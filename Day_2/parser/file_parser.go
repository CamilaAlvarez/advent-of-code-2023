package parser

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Round struct {
	BlueCubes  int
	RedCubes   int
	GreenCubes int
}

type Game struct {
	Id     int
	Rounds []Round
}

func ParseInputFile(input io.Reader) []Game {
	var games []Game
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		var rounds []Round
		splitLine := strings.Split(line, ":")
		if len(splitLine) < 2 {
			fmt.Println("Invalid line: ", line)
			continue
		}
		idString := strings.Trim(strings.ReplaceAll(splitLine[0], gameString, ""), " ")
		id, err := strconv.Atoi(idString)
		if err != nil {
			fmt.Println("Invalid id: ", idString)
			continue
		}
		roundsStrings := strings.Split(splitLine[1], ";")
		for _, v := range roundsStrings {
			roundStrings := strings.Split(v, ",")
			var round Round
			for _, v2 := range roundStrings {
				if strings.Index(v2, BlueCube) > 0 {
					numberCubes, err := parseRoundToInt(BlueCube, v2)
					if err != nil {
						fmt.Println("Invalid number of blue cubes: ", v2)
						continue
					}
					round.BlueCubes = numberCubes
				} else if strings.Index(v2, GreenCube) > 0 {
					numberCubes, err := parseRoundToInt(GreenCube, v2)
					if err != nil {
						fmt.Println("Invalid number of green cubes: ", v2)
						continue
					}
					round.GreenCubes = numberCubes
				} else if strings.Index(v2, RedCube) > 0 {
					numberCubes, err := parseRoundToInt(RedCube, v2)
					if err != nil {
						fmt.Println("Invalid number of red cubes: ", v2)
						continue
					}
					round.RedCubes = numberCubes
				}
			}
			rounds = append(rounds, round)
		}
		games = append(games, Game{id, rounds})
	}
	return games
}
