/*
Reverse the Array
https://www.geeksforgeeks.org/write-a-program-to-reverse-an-array-or-string/
*/

package main

import "fmt"

func reverseTheArray(arr []int) {
	start_index := 0
	end_index := len(arr) - 1
	for start_index < end_index {
		arr[start_index], arr[end_index] = arr[end_index], arr[start_index]
		start_index += 1
		end_index -= 1
	}
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr)
	reverseTheArray(arr)
	fmt.Println(arr)

	arr1 := []int{1, 2, 3}
	fmt.Println(arr1)
	reverseTheArray(arr1)
	fmt.Println(arr1)
}
