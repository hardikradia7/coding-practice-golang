/*
Chocolate Distribution Problem
https://www.geeksforgeeks.org/chocolate-distribution-problem/
*/

package main

import (
	"fmt"
	"math"
	"sort"
)

func distributeChocolate(arr []int, m int) int {
	min_diff := math.MaxInt
	sort.Ints(arr)
	n := len(arr) - m
	for i := 0; i <= n; i++ {
		diff := arr[i+m-1] - arr[i]

		if diff < min_diff {
			min_diff = diff
		}
	}
	return min_diff
}

func main() {
	arr := []int{7, 3, 2, 4, 9, 12, 56}
	fmt.Println(distributeChocolate(arr, 3))

	arr = []int{7, 3, 2, 5, 54, 55, 56}
	fmt.Println(distributeChocolate(arr, 3))
}
