package parser

import (
	"bufio"
	"io"
	"log"
	"strings"

	desertmap "github.com/CamilaAlvarez/advent-of-code-2023/Day_8/desert_map"
)

func ParseToMap(file io.Reader) desertmap.MapFile {
	scanner := bufio.NewScanner(file)
	var mapFile desertmap.MapFile
	mapFile.Map = make(map[desertmap.Node]desertmap.NodeMapping)
	isFirstLine := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if isFirstLine {
			mapFile.Pattern = &desertmap.InstructionPattern{Pattern: line}
			isFirstLine = false
		} else {
			lineSplit := strings.Split(line, "=")
			if len(lineSplit) < 2 {
				log.Fatal("Invalid map entry:", line)
			}
			key := strings.Trim(lineSplit[0], " ")
			value := strings.Replace(lineSplit[1], "(", "", 1)
			value = strings.Replace(value, ")", "", 1)
			value = strings.Trim(value, " ")
			splitValue := strings.Split(value, ",")
			if len(splitValue) < 2 {
				log.Fatal("Invalid value in map entry:", value)
			}
			mapFile.Map[desertmap.Node(key)] = desertmap.NodeMapping{
				L: desertmap.Node(strings.Trim(splitValue[0], " ")),
				R: desertmap.Node(strings.Trim(splitValue[1], " ")),
			}
		}
	}
	return mapFile
}
