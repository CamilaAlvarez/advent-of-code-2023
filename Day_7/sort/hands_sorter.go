package sort

import (
	"sort"

	"github.com/CamilaAlvarez/advent-of-code-2023/Day_7/parser"
)

func SortHands(hands []parser.Hand) {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Strength == hands[j].Strength {
			// We assume that no two hands are the same
			for k := 0; k < parser.HandSize; k++ {
				if hands[i].Cards[k] != hands[j].Cards[k] {
					return hands[i].Cards[k] < hands[j].Cards[k]
				}
			}

		}
		return hands[i].Strength < hands[j].Strength
	})
}
