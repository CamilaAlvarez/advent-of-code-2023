package parser

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const time string = "Time:"
const distance string = "Distance:"

type Race struct {
	Time     int
	Distance int
}

// we're assumming the file follows the requried structure
func Parse(file io.Reader) []Race {
	re := regexp.MustCompile(" +")
	scanner := bufio.NewScanner(file)
	var races []Race
	times := getList(scanner, re, time)
	distances := getList(scanner, re, distance)
	if len(times) != len(distances) {
		log.Fatal("Number of times and distances don't match: ", len(times), " vs. ", len(distances))
	}
	for k, v := range times {
		races = append(races, Race{v, distances[k]})
	}
	return races
}

func getList(scanner *bufio.Scanner, re *regexp.Regexp, remove string) []int {
	var list []int
	if scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, remove, "")
		line = strings.Trim(line, " ")
		for _, v := range re.Split(line, -1) {
			value, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal("Invalid number: ", v)
			}
			list = append(list, value)
		}
	}
	return list
}
