package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getInstancesNumber(2, []int{25, 23, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 76, 80}))
}

func getInstancesNumber(instances int, usages []int) int {
	numInst := float64(instances)
	sleep := 0

	for _, usage := range usages {
		if sleep > 0 {
			sleep--
			continue
		}

		switch {
		case usage < 25:
			if numInst > 1 {
				numInst = math.Ceil(float64(numInst / 2))
				sleep += 10
			}
		case usage > 60:
			doubleVal := numInst * 2.0
			if doubleVal < 2*math.Pow(10.0, 8.0) {
				numInst = doubleVal
				sleep += 10
			}
		default:
			fmt.Println("Nothing to do")
		}
	}

	return int(numInst)
}
