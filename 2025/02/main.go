package main

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"

	"github.com/rodrigoaraujo46/aoc/2025/readers"
	"github.com/rodrigoaraujo46/assert"
)

const todaysPath = "./2025/02"

func main() {
	operations := strings.FieldsFuncSeq(readers.ReadInput(todaysPath), func(r rune) bool {
		return r == ',' || r == '\n'
	})
	/*
		operations := strings.FieldsFuncSeq(readers.ReadInput(todaysPath), func(r rune) bool {
			return r == ',' || r == '\n'
		})
	*/

	/*
		ans := 0
		for op := range operations {
			nums := strings.SplitN(op, "-", 2)
			start, errStart := strconv.Atoi(nums[0])
			end, errEnd := strconv.Atoi(nums[1])
			assert.NoError(cmp.Or(errStart, errEnd), "Couldn't convert range to numbers")

			for ; start <= end; start++ {
				s := strconv.Itoa(start)
				if len(s)%2 != 0 {
					continue
				}
				mid := len(s) / 2
				first := s[:mid]
				second := s[mid:]
				if first == second {
					ans += start
				}
			}
		}

		fmt.Println(ans)
	*/

	ans := 0
	for op := range operations {
		nums := strings.SplitN(op, "-", 2)
		start, errStart := strconv.Atoi(nums[0])
		end, errEnd := strconv.Atoi(nums[1])
		assert.NoError(cmp.Or(errStart, errEnd), "Couldn't convert range to numbers")

		for ; start <= end; start++ {
			s := strconv.Itoa(start)
		sizeLoop:
			for size := len(s) / 2; size > 0; size-- {
				if len(s)%size != 0 {
					continue
				}
				first := s[:size]
				for i := size; i < len(s); i += size {
					if s[i:i+size] != first {
						continue sizeLoop
					}
				}
				ans += start
				break
			}
		}
	}

	fmt.Println(ans)
}
