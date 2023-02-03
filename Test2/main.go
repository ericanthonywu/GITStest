package main

import (
	"fmt"
	"sort"
)

func main() {
	//run([]int{100, 100, 50, 40, 40, 20, 10}, []int{5, 25, 50, 120})
	run([]int{100, 80, 80, 70}, []int{60, 70, 100})
}

func run(skorPemain []int, skorGits []int) {
	// append all array
	totalSkor := append(skorPemain, skorGits...)

	// sort array
	sort.Slice(totalSkor, func(i, j int) bool {
		return totalSkor[i] > totalSkor[j]
	})

	// remove duplicates
	totalSkor = unique(totalSkor)

	for _, gits := range skorGits {
		fmt.Println(getIndexByValue(totalSkor, gits), " ")
	}
}

func getIndexByValue(arr []int, value int) int {
	for i, arrValue := range arr {
		if arrValue == value {
			return i + 1
		}
	}
	panic("value is not found")
}

func unique(arr []int) []int {
	occurred := map[int]bool{}
	var result []int
	for e := range arr {
		if occurred[arr[e]] != true {
			occurred[arr[e]] = true

			result = append(result, arr[e])
		}
	}

	return result
}
