package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_5/parser"
)

func binarySearch(s int, elementMap []parser.ElementMapRow) int {
	if len(elementMap) == 0 {
		// if there's no smaller division to be made (length array == 0)
		// we couldn't find a mapping, so we return the same value
		return s
	}
	halfLength := int(math.Floor(float64(len(elementMap)) / 2))
	middleItem := elementMap[halfLength]
	if s >= middleItem.SourceStart && s < (middleItem.SourceStart+middleItem.RangeLength) {
		// 1. Check if in range,
		return (s - middleItem.SourceStart) + middleItem.DestinationStart
	} else if middleItem.SourceStart > s {
		// 2. If source > s, search in the lower half
		return binarySearch(s, elementMap[:halfLength])

	} else if (middleItem.SourceStart + middleItem.RangeLength) <= s {
		// 3. If source + range  >= s, search in the upper half
		return binarySearch(s, elementMap[halfLength+1:])
	}
	// If we reached here is because there is only one item left in the map, and s doesn't belong to its range
	return s
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Missing arguments: ", len(os.Args))
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Could not open input file: ", os.Args[1])
	}
	lowestLocation := math.MaxInt
	almanac := parser.ParseInputToAlmanac(file)
	for _, s := range almanac.Seeds {
		// Phase 1: Seed to soil
		soil := binarySearch(s, almanac.SeedToSoilMap)
		fmt.Println("Soil for seed ", s, " is ", soil)
		// Phase 2: Soil to fertilizer
		fertilizer := binarySearch(soil, almanac.SoilToFertilizerMap)
		fmt.Println("Fertilizer for seed ", s, " is ", fertilizer)
		// Phase 3: Fertilizer to water
		water := binarySearch(fertilizer, almanac.FertilizerToWaterMap)
		fmt.Println("Water for seed ", s, " is ", water)
		// Phase 4: Water to light
		light := binarySearch(water, almanac.WaterToLightMap)
		fmt.Println("Light for seed ", s, " is ", light)
		// Pase 5: Light to temperature
		temperature := binarySearch(light, almanac.LightToTemperature)
		fmt.Println("Temperature for seed ", s, " is ", temperature)
		// Phase 6: Temperature to humidity
		humidity := binarySearch(temperature, almanac.TemperatureToHumididty)
		fmt.Println("Humidity for seed ", s, " is ", humidity)
		// Phase 7: Humidity to location
		location := binarySearch(humidity, almanac.HumidityToLocation)
		if location < lowestLocation {
			lowestLocation = location
		}
		fmt.Println("Location for seed ", s, " is ", location)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("Lowest location: ", lowestLocation)
}
