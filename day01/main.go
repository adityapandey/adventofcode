package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	sum0, sum1 := 0, 0
	l := len(b)
	for i := range b {
		if b[i] == b[(i+1)%l] {
			sum0 += int(b[i]) - '0'
		}
		if b[i] == b[(i+l/2)%l] {
			sum1 += int(b[i]) - '0'
		}
	}
	fmt.Println(sum0)
	fmt.Println(sum1)
}
