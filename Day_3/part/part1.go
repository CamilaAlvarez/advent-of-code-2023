package part

import (
	"fmt"
	"unicode"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_3/math"
)

func fillCollisionMatrix(collisionMatrix [][]bool, matrix [][]rune) {
	for i, r := range matrix {
		for j, v := range r {
			if !unicode.IsDigit(v) && v != '.' {
				for k := -1; k < 2; k++ {
					collisionI := i + k
					if collisionI < 0 || collisionI > len(matrix)-1 {
						continue
					}
					for l := -1; l < 2; l++ {
						collisionJ := j + l
						if collisionJ < 0 || collisionJ > len(r)-1 || (i == collisionI && collisionJ == j) {
							continue
						}
						if unicode.IsDigit(matrix[collisionI][collisionJ]) {
							collisionMatrix[collisionI][collisionJ] = true
						}
					}
				}
			}
		}
	}
}
func getSchematicsSum(collisionMatrix [][]bool, matrix [][]rune) {
	var sumSchematics int
	for i, r := range matrix {
		var number, powerOfTen int
		var hasCollision bool
		for j := len(r) - 1; j >= 0; j-- {
			currentSymbol := r[j]
			if unicode.IsDigit(currentSymbol) {
				number += (int(currentSymbol-'0') * math.IntPow(10, powerOfTen))
				powerOfTen++
				if collisionMatrix[i][j] {
					hasCollision = true
				}
			} else {
				if hasCollision {
					sumSchematics += number

				}
				hasCollision = false
				powerOfTen = 0
				number = 0
			}
		}
		// for the last iteration
		if hasCollision {
			sumSchematics += number

		}
	}
	fmt.Println("Sum schematics: ", sumSchematics)
}
func Part1(collisionMatrix [][]bool, matrix [][]rune) {
	fillCollisionMatrix(collisionMatrix, matrix)
	getSchematicsSum(collisionMatrix, matrix)
}
