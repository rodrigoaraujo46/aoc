package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rodrigoaraujo46/aoc/2025/readers"
	"github.com/rodrigoaraujo46/assert"
)

const todaysPath = "./2025/01"

func main() {
	// operations := strings.SplitSeq(readers.ReadExample(todaysPath), "\n")
	operations := strings.SplitSeq(readers.ReadInput(todaysPath), "\n")

	/* 	PART 1
	ans, dial := 0, 50
	for operation := range operations {
		if operation == "" {
			continue
		}
		isClockwise := operation[0] == 'R'
		num, err := strconv.Atoi(operation[1:])
		assert.NoError(err, "Couldn't get number from operation")

		dial -= num
		if isClockwise {
			dial += 2 * num
		}
		dial = dial % 100
		if dial == 0 {
			ans++
		}
	}
	*/

	// PART 2
	ans, dial := 0, 50
	for operation := range operations {
		if operation == "" {
			continue
		}
		isClockwise := operation[0] == 'R'
		num, err := strconv.Atoi(operation[1:])
		assert.NoError(err, "Couldn't get number from operation")

		ans += num / 100

		num %= 100
		if isClockwise {
			if dial != 0 && dial+num > 100 {
				ans++
			}
			dial = (dial + num) % 100
		} else {
			if dial != 0 && dial-num < 0 {
				ans++
			}
			dial = (dial - num + 100) % 100
		}

		if dial == 0 {
			ans++
		}
	}

	fmt.Println(ans)
}
