/*
Next Permutation
https://leetcode.com/problems/next-permutation/
*/

package main

import "fmt"

func reverse(nums []int, i, n int) {
	j := n - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1, n)

}

func main() {
	arr := []int{1, 3, 2}
	nextPermutation(arr)
	fmt.Println(arr)
}
