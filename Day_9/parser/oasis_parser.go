package parser

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_9/oasis"
)

func ParseReport(file io.Reader) oasis.Report {
	scanner := bufio.NewScanner(file)
	var report oasis.Report
	for scanner.Scan() {
		var values []int
		line := scanner.Text()
		line = strings.Trim(line, " ")
		if len(line) == 0 {
			continue
		}
		splitValues := strings.Split(line, " ")
		// We're assuming the input is correclty formed and all reports have the same dimension
		for _, v := range splitValues {
			hv, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal("Invalid report value: ", v)
			}
			values = append(values, hv)
		}
		report.Histories = append(report.Histories, oasis.History{Values: values})
	}
	return report
}
