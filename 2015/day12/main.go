package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	b := []byte(util.ReadAll())
	var f interface{}
	if err := json.Unmarshal(b, &f); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sumNum(f, false))
	fmt.Println(sumNum(f, true))
}

func sumNum(f interface{}, ignoreRed bool) float64 {
	var sum float64
	switch ff := f.(type) {
	case float64:
		sum += ff
	case []interface{}:
		for i := range ff {
			sum += sumNum(ff[i], ignoreRed)
		}
	case map[string]interface{}:
		var msum float64
		var red bool
		for _, v := range ff {
			if vv, ok := v.(string); ok && vv == "red" {
				red = true
			}
			msum += sumNum(v, ignoreRed)
		}
		if !(ignoreRed && red) {
			sum += msum
		}
	}
	return sum
}
