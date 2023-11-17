package nnkosmos

import (
	"golang.org/x/exp/slices"
)

// 题目：
// 给定两个数组 a 和 b，求交集，时间复杂度是多少？
// 给定的数组中无重复元素。

type intersectSolution func([]int, []int) []int

// 1. 暴力法 O(a*b)
func intersectBruteForce(a []int, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}

	var ret []int
	for _, num1 := range a {
		for _, num2 := range b {
			if num1 == num2 {
				ret = append(ret, num1)
			}
		}
	}
	return ret
}

// 2. 哈希表 O(a+b)
func intersectMap(a []int, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}

	m := make(map[int]struct{}, len(a))
	for _, ele := range a {
		m[ele] = struct{}{}
	}

	var ret []int
	for _, ele := range b {
		if _, ok := m[ele]; ok {
			ret = append(ret, ele)
		}
	}
	return ret
}

// 3. 排序和双指针
// O(a*loga) + O(b*logb) + O(a) + O(b) => O(a*loga + b*logb)
// 时间主要花在排序上
func intersectSortAndTwoPointer(a []int, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}

	slices.Sort(a)
	slices.Sort(b)

	var (
		i, j int
		ret  []int
	)
	for {
		if a[i] == b[j] {
			ret = append(ret, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
		if i >= len(a) || j >= len(b) {
			break
		}
	}

	return ret
}

// 4. 排序和二分查找
// O(a*loga) + O(b*loga) => O((a+b)*loga)
func intersectSortAndBinarySearch(a []int, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}

	slices.Sort(a)
	var ret []int
	for _, num := range b {
		if _, found := slices.BinarySearch(a, num); found {
			ret = append(ret, num)
		}
	}
	return ret
}
