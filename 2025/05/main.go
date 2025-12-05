package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/rodrigoaraujo46/aoc/2025/readers"
)

const todaysPath = "./2025/05"

func main() {
	lines := strings.Split(readers.ReadInput(todaysPath), "\n")
	lines = lines[:len(lines)-1]

	/*
		i := slices.Index(lines, "")
		ranges, set := lines[:i], lines[i+1:]

		fresh := 0
		for _, ingredientS := range set {
			ingredient, _ := strconv.Atoi(ingredientS)
			for _, freshRange := range ranges {
				rangeSplit := strings.Split(freshRange, "-")
				startS, endS := rangeSplit[0], rangeSplit[1]
				start, _ := strconv.Atoi(startS)
				end, _ := strconv.Atoi(endS)
				if start <= ingredient && ingredient <= end {
					fresh++
					break
				}
			}
		}

		fmt.Println(fresh)
	*/

	i := slices.Index(lines, "")
	ranges := lines[:i]

	parsed := make([][]int, 0, len(ranges))
	for _, rr := range ranges {
		parts := strings.Split(rr, "-")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		if a > b {
			a, b = b, a
		}
		parsed = append(parsed, []int{a, b})
	}

	sort.Slice(parsed, func(i, j int) bool {
		return parsed[i][0] < parsed[j][0]
	})

	merged := [][]int{parsed[0]}
	for _, curr := range parsed[1:] {
		last := merged[len(merged)-1]
		if curr[0] <= last[1] {
			if curr[1] > last[1] {
				last[1] = curr[1]
			}
			continue
		}
		merged = append(merged, curr)
	}

	ans := 0
	for _, r := range merged {
		ans += r[1] - r[0] + 1
	}

	fmt.Println(ans)
}
