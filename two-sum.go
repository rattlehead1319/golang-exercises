package main

/****
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
***/

import (
	"fmt"
)

func main() {
	res := twoSum([]int{3,5,7,2,8,4}, 11)
	for _, x := range res {
		fmt.Println(x)
	}
}


func twoSum(nums []int, target int) []int {
    imap := make(map[int]int)

    for i, x := range nums {
        c := target - x
        if val, ok := imap[c]; ok {
            return []int{val, i}
        }
        imap[x] = i
    }
    return []int{-1, -1}
}