package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rodrigoaraujo46/aoc/2025/readers"
)

const todaysPath = "./2025/03"

func main() {
	operations := strings.SplitSeq(readers.ReadInput(todaysPath), "\n")

	/*
		ans := 0
		for op := range operations {
			if op == "" {
				continue
			}

			maxi := 0
			for i := range len(op) - 1 {
				for j := i + 1; j < len(op); j++ {
					maxi = max(maxi, 10*int(op[i]-'0')+int(op[j]-'0'))
				}
			}

			ans += maxi
		}
		fmt.Println("ans", ans)
	*/

	const k = 12
	ans := 0
	for op := range operations {
		if op == "" {
			continue
		}
		res := make([]byte, 0, k)
		start := 0

		for picks := range k {
			need := k - picks - 1
			maxDigit := byte('0')
			maxIndex := start

			end := len(op) - need
			for i := start; i < end; i++ {
				if op[i] > maxDigit {
					maxDigit = op[i]
					maxIndex = i
				}
			}

			res = append(res, maxDigit)
			start = maxIndex + 1
		}

		num, _ := strconv.Atoi(string(res))
		ans += num
	}

	fmt.Println("ans", ans)
}
