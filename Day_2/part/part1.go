package part

import (
	"fmt"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_2/parser"
)

func Part1(games []parser.Game) {

	redCubes := 12
	blueCubes := 14
	greenCubes := 13
	var validGamesIdsSum int
gameLoop:
	for _, v := range games {
		for _, round := range v.Rounds {
			if round.BlueCubes > blueCubes || round.GreenCubes > greenCubes || round.RedCubes > redCubes {
				fmt.Printf(`
Invalid number of cubes: 
	Round red %d vs available red %d
	Round blue %d vs available blue %d
	Round green %d vs available green %d
					
	`, round.RedCubes, redCubes, round.BlueCubes, blueCubes, round.GreenCubes, greenCubes)
				continue gameLoop
			}
		}
		validGamesIdsSum += v.Id
	}
	fmt.Println("Sum ids: ", validGamesIdsSum)
}
