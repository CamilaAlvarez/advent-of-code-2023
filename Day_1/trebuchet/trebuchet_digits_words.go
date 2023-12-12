package trebuchet

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func Part2() {
	numbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	if len(os.Args) <= 1 {
		log.Fatal("Missing some arguments. Current number: ", len(os.Args))
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Could not open file: ", os.Args[1])
	}
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for k, v := range numbers {
			// Add the number like this: <word><number><word>, that way we can keep shared letters with following numbers
			line = strings.ReplaceAll(line, k, k+string(v)+k)
		}
		digits := make([]int, 0, len(line))
		for _, v := range line {
			if unicode.IsDigit(v) {
				digits = append(digits, int(v-'0'))
			}
		}
		if len(digits) <= 0 {
			fmt.Println("Line does not contain any digits: ", line)
			continue
		}
		firstDigit := digits[0]
		lastDigit := digits[len(digits)-1]
		sum += (firstDigit*10 + lastDigit)
	}
	fmt.Printf("Sum of calibration values: %d\n", sum)
}
