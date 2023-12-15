package parser

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const seeds string = "seeds:"
const seedToSoil string = "seed-to-soil map:"
const soilToFertilizer string = "soil-to-fertilizer map:"
const fertilizerToWater string = "fertilizer-to-water map:"
const waterToLight string = "water-to-light map:"
const lightToTemperature string = "light-to-temperature map:"
const temperatureToHumidity string = "temperature-to-humidity map:"
const humidityToLocation string = "humidity-to-location map:"

type ElementMapRow struct {
	DestinationStart int
	SourceStart      int
	RangeLength      int
}
type SeedRange struct {
	SeedStart       int
	SeedRangeLength int
}

type Almanac struct {
	Seeds                  []SeedRange
	SeedToSoilMap          []ElementMapRow
	SoilToFertilizerMap    []ElementMapRow
	FertilizerToWaterMap   []ElementMapRow
	WaterToLightMap        []ElementMapRow
	LightToTemperature     []ElementMapRow
	TemperatureToHumididty []ElementMapRow
	HumidityToLocation     []ElementMapRow
}

func ParseInputToAlmanac(file io.Reader) Almanac {
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(" +")
	var almanac Almanac
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if strings.Index(line, seeds) == 0 {
			line = strings.Replace(line, seeds, "", 1)
			seeds := createNumberListFromString(re, line)
			if len(seeds)%2 != 0 {
				log.Fatal("Invalid number of seeds (should always be even)")
			}
			var seedRanges []SeedRange
			for i := 0; i < len(seeds)/2; i++ {
				seedRanges = append(seedRanges, SeedRange{seeds[i*2], seeds[i*2+1]})
			}
			almanac.Seeds = seedRanges
		} else {
			switch line {
			case seedToSoil:
				elementMap := createRangeRows(scanner, re)
				almanac.SeedToSoilMap = elementMap
			case soilToFertilizer:
				elementMap := createRangeRows(scanner, re)
				almanac.SoilToFertilizerMap = elementMap
			case fertilizerToWater:
				elementMap := createRangeRows(scanner, re)
				almanac.FertilizerToWaterMap = elementMap
			case waterToLight:
				elementMap := createRangeRows(scanner, re)
				almanac.WaterToLightMap = elementMap
			case lightToTemperature:
				elementMap := createRangeRows(scanner, re)
				almanac.LightToTemperature = elementMap
			case temperatureToHumidity:
				elementMap := createRangeRows(scanner, re)
				almanac.TemperatureToHumididty = elementMap
			case humidityToLocation:
				elementMap := createRangeRows(scanner, re)
				almanac.HumidityToLocation = elementMap
			default:
				fmt.Println("Invalid transformation: ", line)
			}
		}
	}
	return almanac
}

func createRangeRows(scanner *bufio.Scanner, re *regexp.Regexp) []ElementMapRow {
	var elementMap []ElementMapRow

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		mapRow := createNumberListFromString(re, line)
		if len(mapRow) != 3 {
			fmt.Println("Invalid rule: ", line)
			continue
		}
		elementMap = append(elementMap, ElementMapRow{mapRow[0], mapRow[1], mapRow[2]})
	}
	sort.Slice(elementMap, func(i, j int) bool {
		return elementMap[i].SourceStart < elementMap[j].SourceStart
	})
	return elementMap
}

func createNumberListFromString(re *regexp.Regexp, line string) []int {
	line = strings.Trim(line, " ")
	values := re.Split(line, -1)
	list := make([]int, 0, len(values))
	for _, v := range values {
		iv, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Invalid seed value: ", v)
			continue
		}
		list = append(list, iv)
	}
	return list
}
