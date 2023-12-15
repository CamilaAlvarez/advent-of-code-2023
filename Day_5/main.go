package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_5/parser"
)

func binarySearch(s parser.SeedRange, elementMap []parser.ElementMapRow) []parser.SeedRange {
	if len(elementMap) == 0 {
		// if there's no smaller division to be made (length array == 0)
		// we couldn't find a mapping, so we return the same value
		return []parser.SeedRange{s}
	}
	halfLength := int(math.Floor(float64(len(elementMap)) / 2))
	middleItem := elementMap[halfLength]
	startSeed := s.SeedStart
	endSeed := s.SeedStart + s.SeedRangeLength
	if startSeed >= middleItem.SourceStart && startSeed < (middleItem.SourceStart+middleItem.RangeLength) && endSeed <= (middleItem.SourceStart+middleItem.RangeLength) {
		// 1. Check if in range, the seed range is completely covered by the row range
		startValue := (s.SeedStart - middleItem.SourceStart) + middleItem.DestinationStart
		return []parser.SeedRange{{SeedStart: startValue, SeedRangeLength: s.SeedRangeLength}}

	} else if startSeed >= middleItem.SourceStart && startSeed < (middleItem.SourceStart+middleItem.RangeLength) && endSeed > (middleItem.SourceStart+middleItem.RangeLength) {
		// 2. Check if in range, the seed range is not completely covered by the row range
		// we split the range
		startValue := (s.SeedStart - middleItem.SourceStart) + middleItem.DestinationStart
		outsideRange := s.SeedStart + s.SeedRangeLength - (middleItem.SourceStart + middleItem.RangeLength)
		parsedRanges := []parser.SeedRange{{SeedStart: startValue, SeedRangeLength: s.SeedRangeLength - outsideRange}}
		unparsedRange := parser.SeedRange{SeedStart: endSeed - outsideRange, SeedRangeLength: outsideRange}
		// Since we know that this element couldn't fit the range we continue looking in the upper segments
		// (we know that the top part of the range didn't fit, so it cannot be in lower rows)
		parsedRanges = append(parsedRanges, binarySearch(unparsedRange, elementMap[halfLength+1:])...)
		return parsedRanges

	} else if startSeed < middleItem.SourceStart && endSeed >= middleItem.SourceStart && endSeed <= (middleItem.SourceStart+middleItem.RangeLength) {
		// 3. The range of the seed and the row partially match in the lower half
		// we split the range
		var parsedRanges []parser.SeedRange
		unparsedLength := (middleItem.SourceStart - startSeed)
		unparsedRange := parser.SeedRange{SeedStart: s.SeedStart, SeedRangeLength: unparsedLength}
		parsedRanges = append(parsedRanges, binarySearch(unparsedRange, elementMap[:halfLength])...)
		parsedRanges = append(parsedRanges, parser.SeedRange{SeedStart: middleItem.DestinationStart, SeedRangeLength: s.SeedRangeLength - unparsedLength})
		return parsedRanges

	} else if startSeed < middleItem.SourceStart && endSeed > (middleItem.SourceStart+middleItem.RangeLength) {
		// 4. The seed range is larger than the size of the range of the row
		var parsedRanges []parser.SeedRange
		unparsedLowerRange := parser.SeedRange{SeedStart: s.SeedStart, SeedRangeLength: middleItem.SourceStart - s.SeedStart}
		parsedRanges = append(parsedRanges, binarySearch(unparsedLowerRange, elementMap[:halfLength])...)
		parsedRanges = append(parsedRanges, parser.SeedRange{SeedStart: middleItem.DestinationStart, SeedRangeLength: middleItem.RangeLength})
		endRowRange := middleItem.SourceStart + middleItem.RangeLength
		unparsedUpperRange := parser.SeedRange{SeedStart: endRowRange, SeedRangeLength: endSeed - endRowRange}
		parsedRanges = append(parsedRanges, binarySearch(unparsedUpperRange, elementMap[halfLength+1:])...)
		return parsedRanges

	} else if middleItem.SourceStart > startSeed && middleItem.SourceStart > endSeed {
		// 5. If source > s, search in the lower half
		return binarySearch(s, elementMap[:halfLength])

	} else if (middleItem.SourceStart + middleItem.RangeLength) <= startSeed {
		// 6. If source + range  >= s, search in the upper half
		// Note that if the start doesn't belong to this range, then no item in the range will find a spot
		// (the seed range is a monotonically increasing function)
		return binarySearch(s, elementMap[halfLength+1:])
	}
	// If we reached here is because there is only one item left in the map, and s doesn't belong to its range
	return []parser.SeedRange{s}
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
		var fertilizers []parser.SeedRange
		for _, v := range soil {
			fertilizers = append(fertilizers, binarySearch(v, almanac.SoilToFertilizerMap)...)
		}
		fmt.Println("Fertilizer for seed ", s, " is ", fertilizers)

		// Phase 3: Fertilizer to water
		var waters []parser.SeedRange
		for _, v := range fertilizers {
			waters = append(waters, binarySearch(v, almanac.FertilizerToWaterMap)...)
		}
		fmt.Println("Water for seed ", s, " is ", waters)

		// Phase 4: Water to light
		var lights []parser.SeedRange
		for _, v := range waters {
			lights = append(lights, binarySearch(v, almanac.WaterToLightMap)...)
		}
		fmt.Println("Light for seed ", s, " is ", lights)

		// Pase 5: Light to temperature
		var temperatures []parser.SeedRange
		for _, v := range lights {
			temperatures = append(temperatures, binarySearch(v, almanac.LightToTemperature)...)
		}
		fmt.Println("Temperature for seed ", s, " is ", temperatures)

		// Phase 6: Temperature to humidity
		var humidities []parser.SeedRange
		for _, v := range temperatures {
			humidities = append(humidities, binarySearch(v, almanac.TemperatureToHumididty)...)
		}
		fmt.Println("Humidity for seed ", s, " is ", humidities)

		// Phase 7: Humidity to location
		var locations []parser.SeedRange
		for _, v := range humidities {
			locations = append(locations, binarySearch(v, almanac.HumidityToLocation)...)
		}
		for _, l := range locations {
			if l.SeedStart < lowestLocation {
				lowestLocation = l.SeedStart
			}
		}
		fmt.Println("Location for seed ", s, " is ", locations)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("Lowest location: ", lowestLocation)
}
