package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_2/parser"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Missing arguments: ", len(os.Args))
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Could not open input file: ", os.Args[1])
	}
	redCubes, err := parser.ParseCubeToNumber(os.Args[2], parser.RedCube)
	if err != nil {
		log.Fatal("Invalid number of red cubes: ", os.Args[2])
	}
	blueCubes, err := parser.ParseCubeToNumber(os.Args[3], parser.BlueCube)
	if err != nil {
		log.Fatal("Invalid number of blue cubes: ", os.Args[3])
	}
	greenCubes, err := parser.ParseCubeToNumber(os.Args[4], parser.GreenCube)
	if err != nil {
		log.Fatal("Invalid number of green cubes: ", os.Args[4])
	}
	games := parser.ParseInputFile(file)
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
