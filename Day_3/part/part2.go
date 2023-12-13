package part

import (
	"fmt"
	"unicode"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_3/math"
)

const gear rune = '*'

func fillGearCollisionList(matrix [][]rune) [][2]int {
	var collisionList [][2]int
	for i, r := range matrix {
		for j, v := range r {
			if v == gear {
				var collisionsWithNumbers int
				for k := -1; k < 2; k++ {
					collisionI := i + k
					latestCollision := -1
					if collisionI < 0 || collisionI > len(matrix)-1 {
						continue
					}
					for l := -1; l < 2; l++ {
						collisionJ := j + l
						if collisionJ < 0 || collisionJ > len(r)-1 || (i == collisionI && collisionJ == j) {
							continue
						}
						if unicode.IsDigit(matrix[collisionI][collisionJ]) {
							// we check that we're not considering the same number twice
							if latestCollision != collisionJ-1 && latestCollision != collisionJ+1 {
								collisionsWithNumbers++
							}
							latestCollision = collisionJ

						}
					}
				}
				// We only mark those number that collide with gears that, at the same time, collide with only two number
				if collisionsWithNumbers == 2 {
					collisionList = append(collisionList, [2]int{i, j})
				}
			}
		}
	}
	return collisionList
}
func generateProcessedMatrix(matrix [][]rune) [][]int {
	processedMatrix := make([][]int, len(matrix))
	for i, r := range matrix {
		processedMatrix[i] = make([]int, len(r))
		var number, powerOfTen int
		var initialJ int
		var hasSetInitialJ bool
		for j := len(r) - 1; j >= 0; j-- {
			currentSymbol := r[j]
			if unicode.IsDigit(currentSymbol) {
				if !hasSetInitialJ {
					hasSetInitialJ = true
					initialJ = j
				}
				number += (int(currentSymbol-'0') * math.IntPow(10, powerOfTen))
				powerOfTen++
			} else {
				if hasSetInitialJ {
					for k := initialJ; k > j; k-- {
						processedMatrix[i][k] = number
					}
				}
				number = 0
				powerOfTen = 0
				hasSetInitialJ = false
			}
		}
		if hasSetInitialJ {
			for k := initialJ; k >= 0; k-- {
				processedMatrix[i][k] = number
			}
		}
	}
	return processedMatrix
}
func getGearSchematicsSum(collisionList [][2]int, matrix [][]rune) {
	var sumSchematics int
	processedMatrix := generateProcessedMatrix(matrix)
	for _, v := range collisionList {
		i := v[0]
		j := v[1]
		multiplication := 1
		for k := -1; k < 2; k++ {
			collisionI := i + k
			latestCollision := -1
			if collisionI < 0 || collisionI > len(matrix)-1 {
				continue
			}
			for l := -1; l < 2; l++ {
				collisionJ := j + l
				if collisionJ < 0 || collisionJ > len(matrix[0])-1 || (i == collisionI && collisionJ == j) {
					continue
				}
				if processedMatrix[collisionI][collisionJ] != 0 {
					// we check that we're not considering the same number twice
					if latestCollision != collisionJ-1 && latestCollision != collisionJ+1 {
						multiplication *= processedMatrix[collisionI][collisionJ]
					}
					latestCollision = collisionJ

				}
			}
		}
		sumSchematics += multiplication
	}
	fmt.Println("Sum schematics: ", sumSchematics)
}
func Part2(matrix [][]rune) {
	collisionList := fillGearCollisionList(matrix)
	getGearSchematicsSum(collisionList, matrix)
}
