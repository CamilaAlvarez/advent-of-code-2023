package parser

import (
	"bufio"
	"io"
	"log"

	pipemap "github.com/CamilaAlvarez/advent-of-code-2023/Day_10/pipe_map"
)

func ParseMap(file io.Reader) pipemap.PipeMap {
	var pipes pipemap.PipeMap
	scanner := bufio.NewScanner(file)
	var rowNumberStart int
	for scanner.Scan() {
		var colNumberStart int
		line := scanner.Text()
		var row []pipemap.Pipe
		for _, c := range line {
			switch c {
			case ' ':
				continue
			case 'S':
				row = append(row, pipemap.S)
				pipes.StartLocationRow = rowNumberStart
				pipes.StartLocationCol = colNumberStart
			case '|':
				row = append(row, pipemap.I)
			case '-':
				row = append(row, pipemap.H)
			case 'L':
				row = append(row, pipemap.L)
			case 'J':
				row = append(row, pipemap.J)
			case '7':
				row = append(row, pipemap.Seven)
			case 'F':
				row = append(row, pipemap.F)
			case '.':
				row = append(row, pipemap.Dot)
			default:
				log.Fatal("Invalid symbol: ", c)
			}
			colNumberStart++
		}
		pipes.Map = append(pipes.Map, row)
		rowNumberStart++
	}
	return pipes
}
