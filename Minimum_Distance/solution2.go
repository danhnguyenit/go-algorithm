package main

import (
	"fmt"
	"math"
	"sort"
)

func minimunDistance(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	smallest, cur := math.MaxInt32, math.MaxInt32
	var idxOne, idxTwo int
	var pairNumber []int
	for idxOne < len(array1) && idxTwo < len(array2) {
		first, second := array1[idxOne], array2[idxTwo]
		if first == second {
			return []int{first, second}
		}
		if first < second {
			idxOne += 1
			cur = second - first

		} else {
			idxTwo += 1
			cur = first - second
		}

		if cur < smallest {
			smallest = cur
			pairNumber = []int{first, second}
		}
	}
	return pairNumber
}
func main() {
	array1 := []int{-1, 5, 10, 20, 28, 3}
	array2 := []int{26, 134, 135, 15, 17}
	fmt.Println(minimunDistance(array1, array2))
}
