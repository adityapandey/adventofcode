package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type num struct {
	pos, val int
}

func main() {
	var nums []num
	for i, n := range strings.Split(util.ReadAll(), "\n") {
		nums = append(nums, num{i, util.Atoi(n)})
	}
	nums2 := make([]num, len(nums))
	for i := range nums {
		nums2[i] = num{nums[i].pos, 811589153 * nums[i].val}
	}

	mix(nums)
	fmt.Println(coords(nums))

	for i := 0; i < 10; i++ {
		mix(nums2)
	}
	fmt.Println(coords(nums2))
}

func mix(nums []num) {
	n := len(nums) - 1
	for i := range nums {
		oldpos := nums[i].pos
		newpos := ((oldpos+nums[i].val)%n + n) % n
		if oldpos < newpos {
			for j := range nums {
				if nums[j].pos > oldpos && nums[j].pos <= newpos {
					nums[j].pos--
				}
			}
		}
		if newpos < oldpos {
			for j := range nums {
				if nums[j].pos >= newpos && nums[j].pos < oldpos {
					nums[j].pos++
				}
			}
		}
		nums[i].pos = newpos
	}
}

func coords(nums []num) int {
	l := len(nums)
	var zeroPos int
	for i := range nums {
		if nums[i].val == 0 {
			zeroPos = nums[i].pos
			break
		}
	}
	sum := 0
	for i := range nums {
		if nums[i].pos == (zeroPos+1000)%l || nums[i].pos == (zeroPos+2000)%l || nums[i].pos == (zeroPos+3000)%l {
			sum += nums[i].val
		}
	}
	return sum
}
