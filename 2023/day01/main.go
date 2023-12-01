package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	nums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	val := map[string]int{}
	for i := 1; i <= 9; i++ {
		val[fmt.Sprint(i)] = i
		val[nums[i-1]] = i
	}
	re1 := regexp.MustCompile("[1-9]")
	re2 := regexp.MustCompile("[1-9]|" + strings.Join(nums, "|"))
	for i := range nums {
		nums[i] = rev(nums[i])
	}
	for i := 1; i <= 9; i++ {
		val[nums[i-1]] = i
	}
	re2_rev := regexp.MustCompile("[1-9]|" + strings.Join(nums, "|"))

	var sum1, sum2 int
	s := util.ScanAll()
	for s.Scan() {
		t1 := re1.FindAllString(s.Text(), -1)
		sum1 += 10*val[t1[0]] + val[t1[len(t1)-1]]

		t2 := re2.FindString(s.Text())
		t2_rev := re2_rev.FindString(rev(s.Text()))
		sum2 += 10*val[t2] + val[t2_rev]
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func rev(s string) string {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
