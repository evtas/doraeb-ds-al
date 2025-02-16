package main

import (
	"fmt"
	"sort"
)

// 二分查找
// 时间复杂度O(logn)
// 注意点：
// 1.判断条件是 low <= high
// 2.mid=(low+high)/2这种写法是有问题的，可能会溢出，应该改成low+(high-low)/2 或者 low+(high-low)>>1
// 只能用在顺序表存储的数据结构上，针对的是有序数据
// 场景：插入删除操作不频繁、可以一次排序多次查询
func BiSearch(datas []int, val int) int {
	sort.Ints(datas)
	low, high := 0, len(datas)-1
	for low <= high {
		mid := (low + high) / 2
		if datas[mid] == val {
			return mid
		} else if datas[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func BiSearchRecursion(datas []int, val int) int {
	sort.Ints(datas)
	fmt.Printf("datas: %v\n", datas)
	return recursionSearch(datas, 0, len(datas)-1, val)
}

func recursionSearch(datas []int, low, high int, val int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)>>1
	fmt.Printf("mid: %v midVal: %v val: %v\n", mid, datas[mid], val)
	if datas[mid] == val {
		return mid
	} else if datas[mid] < val {
		return recursionSearch(datas, mid+1, high, val)
	} else {
		return recursionSearch(datas, low, mid-1, val)
	}
}

// 查找第一个值等于给定值的元素
func BiSearchFirstEqual(datas []int, val int) int {
	sort.Ints(datas)
	low, high := 0, len(datas)-1
	for low <= high {
		mid := low + (high-low)>>1
		if datas[mid] == val && (mid == 0 || datas[mid-1] != val) {
			return mid
		} else if datas[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// 查找最后一个值等于给定值的元素
func BiSearchLastEqual(datas []int, val int) int {
	sort.Ints(datas)
	low, high := 0, len(datas)-1
	for low <= high {
		mid := low + (high-low)>>1
		if datas[mid] == val && (mid == len(datas)-1 || datas[mid+1] != val) {
			return mid
		} else if datas[mid] <= val {
			// 等号这里是个坑，有无等号需要看你要找的是第一个还是最后一个
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// 查找第一个大于等于给定值的元素
func BiSearchFirstLarger(datas []int, val int) int {
	sort.Ints(datas)
	low, high := 0, len(datas)-1
	for low <= high {
		mid := low + (high-low)>>1
		if datas[mid] >= val && (mid == 0 || datas[mid-1] < val) {
			return mid
		} else if datas[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// 查找最后一个小于等于给定值的元素
func BiSearchLastLess(datas []int, val int) int {
	sort.Ints(datas)
	low, high := 0, len(datas)-1
	for low <= high {
		mid := low + (high-low)>>1
		if datas[mid] <= val && (mid == len(datas)-1 || datas[mid+1] > val) {
			return mid
		} else if datas[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
