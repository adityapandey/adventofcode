package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode2017-go/util"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	sum0, sum1 := 0, 0
	for s.Scan() {
		line := strings.Split(s.Text(), "\t")
		var nums []int
		for _, n := range line {
			nums = append(nums, util.Atoi(n))
		}
		sort.Ints(nums)
		l := len(nums)
		sum0 += nums[l-1] - nums[0]
		for i := 0; i < l-1; i++ {
			for j := i + 1; j < l; j++ {
				if nums[j]%nums[i] == 0 {
					sum1 += nums[j] / nums[i]
				}
			}
		}
	}
	fmt.Println(sum0)
	fmt.Println(sum1)
}
