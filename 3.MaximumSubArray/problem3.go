/*
Maximum-Subarray
https://leetcode.com/problems/maximum-subarray/
*/

package main

func maxSubArray(nums []int) int {
	max_sum := nums[0]
	n := len(nums)
	if n == 1 {
		return max_sum
	}
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1]+nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max_sum {
			max_sum = nums[i]
		}
	}
	return max_sum
}

func main() {

}
