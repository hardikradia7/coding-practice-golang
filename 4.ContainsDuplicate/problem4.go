/*
Contains Duplicate
https://leetcode.com/problems/contains-duplicate/
*/

package main

import "fmt"

func containsDuplicate(nums []int) bool {
	elem_dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := elem_dict[nums[i]]; ok {
			return true
		} else {
			elem_dict[nums[i]] = 1
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 1, 4}
	fmt.Println(containsDuplicate(arr))
}
