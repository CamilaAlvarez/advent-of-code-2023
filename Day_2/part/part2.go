package part

import (
	"fmt"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_2/parser"
)

func Part2(games []parser.Game) {
	var powerSum int
	for _, v := range games {
		var maxRed, maxBlue, maxGreen int
		for _, r := range v.Rounds {
			if maxRed < r.RedCubes {
				maxRed = r.RedCubes
			}
			if maxBlue < r.BlueCubes {
				maxBlue = r.BlueCubes
			}
			if maxGreen < r.GreenCubes {
				maxGreen = r.GreenCubes
			}
		}
		power := (maxRed * maxBlue * maxGreen)
		powerSum += power
	}
	fmt.Println("Power sum: ", powerSum)
}
