/*
Search an Element in a Sorted and Pivoted Array
https://www.geeksforgeeks.org/search-an-element-in-a-sorted-and-pivoted-array/
*/

package main

import "fmt"

func findElem(arr []int, elem, start, end int) int {
	mid := (start + end) / 2

	if start >= end {
		return -1
	}

	if arr[mid] == elem {
		return mid
	}
	if arr[start] < arr[mid] && arr[start] <= elem && arr[mid] >= elem {
		if arr[start] == elem {
			return start
		} else if arr[mid] == elem {
			return end
		}
		return findElem(arr, elem, start, mid)
	} else {
		return findElem(arr, elem, mid+1, end)
	}

}

func main() {
	arr := []int{5, 6, 7, 8, 9, 10, 1, 2, 3}
	fmt.Println(findElem(arr, 3, 0, len(arr)))

	arr = []int{5, 6, 7, 8, 9, 10, 1, 2, 3}
	fmt.Println(findElem(arr, 30, 0, len(arr)))

	arr = []int{30, 40, 50, 10, 20}
	fmt.Println(findElem(arr, 10, 0, len(arr)))
}
