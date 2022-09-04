/*
Repeat and Missing Number Array
https://www.interviewbit.com/problems/repeat-and-missing-number-array/
*/

package main

import "fmt"

func findDuplicates(arr []int) (int, int) {
	n := len(arr)
	i := 0
	for i < n {
		if arr[i] != i+1 {
			val := arr[i]
			if arr[val-1] == val {
				i++
			} else {
				arr[i], arr[val-1] = arr[val-1], arr[i]
			}
		} else {
			i++
		}
	}
	for i = 0; i < n; i++ {
		if arr[i] != i+1 {
			return arr[i], i + 1
		}
	}

	return 0, 0
}

func main() {
	arr := []int{3, 1, 2, 5, 3}
	a, b := findDuplicates(arr)
	fmt.Println(a, b)
}
