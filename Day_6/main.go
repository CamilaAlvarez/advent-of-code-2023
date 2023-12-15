package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_6/parser"
)

// The race function is f(t) = t*(t'-t), where t' is the total time
// We need all t such that f(t) < record, meaning:
//
//	t^2 - t't + x < 0
//	x1 = (t' + sqrt(t'^2 - 4x))/2
//	x2 = (t' - sqrt(t'^2 - 4x))/2
//
// We know that x2 < x1, so the ts that match the requirements are such that:
//
//	x2 < t < x1
func computeRaceFunctionRoots(race parser.Race) (x1 float64, x2 float64) {
	rootPartCuad := math.Sqrt(math.Pow(float64(race.Time), 2) - 4*float64(race.Distance))
	x1 = (float64(race.Time) + rootPartCuad) / 2
	x2 = (float64(race.Time) - rootPartCuad) / 2
	return x1, x2
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Missing arguments: ", len(os.Args))
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Could not open input file: ", os.Args[1])
	}
	races := parser.Parse(file)
	fmt.Println(races)
	fmt.Println()
	fmt.Println()

	mulOptions := 1
	for _, v := range races {
		x1, x2 := computeRaceFunctionRoots(v)
		options := int(math.Ceil(x1-1)) - int(math.Floor(x2+1)) + 1
		fmt.Println("Race with time", v.Time, "and distance", v.Distance, ":", options, "ways to win")
		mulOptions *= options
	}
	fmt.Println()
	fmt.Println("Multiplication of ways to win: ", mulOptions)
}
