package main

import (
	"fmt"
)

func twoSumIndex(nums []int, target int) []int {
	ret := make([]int, 2, 2)

	for i := 0; i < len(nums); i = i + 1 {
		for j := i + 1; j < len(nums); j = j + 1 {
			if nums[i]+nums[j] == target {
				ret[0] = i
				ret[1] = j
				break
			}
		}
	}
	return ret
}

/*
func twoSumIndexV2(num []int,target int)[]int{
 	ret := make([]int,2)
 	for i,j := range nums{
 		index,ok := nums[target - ]
	}
}
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	res := []int{}
	for index, value := range nums {
		i, ok := m[target-value]
		if ok {
			res = []int{i, index}
			return res
		} else {
			m[value] = index
		}
	}
	return res
}

func twoSumV2(nums []int, target int) []int {
	m := make(map[int]int)
	ret := make([]int, 0)
	for index, value := range nums {
		if i, ok := m[target-value]; ok {
			//ret = []int{index,i}
			ret = []int{i, index}
		} else {
			m[value] = index
		}

	}
	return ret
}

func main() {
	//x :=[]int{0,7,4,3,5,9}
	//x :=[]int{3,2,4}
	x1 := []int{2, 5, 5, 11}
	ret := twoSumIndex(x1, 10)
	fmt.Println("index: ", ret[0], " ", ret[1])
	x := []int{0, 7, 4, 3, 5, 9}
	ret = twoSumIndex(x, 9)
	fmt.Println("index: ", ret[0], " ", ret[1])
	ret = twoSum(x, 9)
	fmt.Println("index: ", ret[0], " ", ret[1])
	ret = twoSumV2(x1, 10)
	fmt.Println("index: ", ret[0], " ", ret[1])
}
