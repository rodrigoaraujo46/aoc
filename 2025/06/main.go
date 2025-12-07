package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/rodrigoaraujo46/aoc/2025/readers"
)

const todaysPath = "./2025/06"

func main() {
	lines := strings.Split(readers.ReadInput(todaysPath), "\n")
	// rows, cols := len(lines), len(strings.Fields(lines[0]))

	/*
		problems := make([][]string, cols)
		for i := range problems {
			problems[i] = make([]string, 0, rows)
		}

		for _, row := range lines {
			x := 0
			for col := range strings.FieldsSeq(row) {
				problems[x] = append(problems[x], col)
				x++
			}
		}

		ans := 0
		for _, problem := range problems {
			sign := problem[len(problem)-1]
			res := 0
			if sign == "*" {
				res = 1
			}
			for _, num := range problem[:len(problem)-1] {
				numInt, _ := strconv.Atoi(num)
				if sign == "*" {
					res *= numInt
					continue
				}
				res += numInt
			}
			ans += res
		}
	*/

	charsFull := make([][]string, 0)
	for _, line := range lines {
		charsFull = append(charsFull, strings.Split(line, ""))
	}

	t := transpose(charsFull)
	t = slices.DeleteFunc(t, func(a []string) bool {
		return !slices.ContainsFunc(a, func(s string) bool {
			return s != " "
		})
	})

	ans := 0
	cur := 0
	sign := ""
	for i := range t {
		if isStart(t[i]) {
			ans += cur
			cur, sign = 0, t[i][len(t[i])-1]
			if sign == "*" {
				cur = 1
			}
		}

		if sign == "*" {
			cur *= toNum(t[i])
		} else {
			cur += toNum(t[i])
		}
	}

	fmt.Println(ans + cur)
}

func isStart(s []string) bool {
	return slices.ContainsFunc(s, func(e string) bool {
		return e == "*" || e == "+"
	})
}

func transpose(matrix [][]string) [][]string {
	maxCols := 0
	for _, row := range matrix {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}
	transposed := make([][]string, maxCols)
	for i := range matrix {
		for j := 0; j < len(matrix[i]); j++ {
			value := matrix[i][j]
			transposed[j] = append(transposed[j], value)
		}
	}

	return transposed
}

func toNum(s []string) int {
	numS := ""
	for _, v := range s {
		if v != " " && v != "*" && v != "+" {
			numS += v
		}
	}

	num, _ := strconv.Atoi(numS)
	return num
}
