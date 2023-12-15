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
func Parse(file io.Reader) Race {
	re := regexp.MustCompile(" +")
	scanner := bufio.NewScanner(file)
	time := getNumberFromLine(scanner, re, time)
	distance := getNumberFromLine(scanner, re, distance)

	return Race{time, distance}
}

func getNumberFromLine(scanner *bufio.Scanner, re *regexp.Regexp, remove string) int {
	var number int
	if scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, remove, "")
		line = strings.Trim(line, " ")
		line = re.ReplaceAllString(line, "")
		value, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("Invalid number: ", line)
		}
		number = value
	}
	return number
}
