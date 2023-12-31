package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_11/parser"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Missing arguments: ", len(os.Args))
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Could not open input file: ", os.Args[1])
	}
	galaxies := parser.ParseToGalaxies(file)
	var sumDistances int
	for i := 0; i < len(galaxies.GalaxyLocation)-1; i++ {
		for j := i + 1; j < len(galaxies.GalaxyLocation); j++ {
			distI := int(math.Abs(float64(galaxies.GalaxyLocation[j].I - galaxies.GalaxyLocation[i].I)))
			distJ := int(math.Abs(float64(galaxies.GalaxyLocation[j].J - galaxies.GalaxyLocation[i].J)))
			sumDistances += (distI + distJ)
		}
	}
	fmt.Println("Sum distances: ", sumDistances)
}
