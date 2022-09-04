/*
Maximum and Minimum Element in an Array
https://www.geeksforgeeks.org/maximum-and-minimum-in-an-array/
*/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	max_elem := arr[0]
	min_elem := arr[1]
	for _, elem := range arr[1:] {
		if elem > max_elem {
			max_elem = elem
		} else if elem < min_elem {
			min_elem = elem
		}
	}
	fmt.Println(max_elem)
	fmt.Println(min_elem)
}
