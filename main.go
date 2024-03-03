package main

import (
	"fmt"
	"sort"
)

func mergeIntervals(intervals [][]int) {
	fmt.Println("Original - ", intervals)
	if len(intervals) == 0 {
		fmt.Println(nil)
		return
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		lastItem := merged[len(merged)-1]
		if lastItem[1] > intervals[i][0] {
			lastItem[1] = intervals[i][1]
		} else {
			merged = append(merged, intervals[i])
		}
	}
	fmt.Println("Final - ", merged)
}

func main() {
	intervals := [][]int{{6, 9}, {7, 10}, {1, 3}, {2, 5}, {11, 13}}
	mergeIntervals(intervals)
}
