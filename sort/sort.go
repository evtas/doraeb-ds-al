package main

import (
	"fmt"
	"sort"
)

// BubbleSort 冒泡排序
// 稳定的原地排序，时间复杂度最好是O(n)，最坏是O(n^2)，平均时间复杂度可以通过有序度、无序度进行粗略估计
func BubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		hasSwap := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j+1] < nums[j] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
				hasSwap = true
			}
		}
		if !hasSwap {
			break
		}
	}
}

// InsertSort 插入排序
// 稳定的原地排序算法，最坏时间复杂度O(n^2)，最好时间复杂度O(n)，平均时间复杂度O(n^2)
func InsertSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	for i := 1; i < len(nums); i++ {
		val := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > val {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = val
	}
}

// ShellSort 希尔排序
func ShellSort(nums []int) {

}

// SelectSort 选择排序
// 非稳定的原地排序算法，最坏时间复杂度O(n^2)，最好时间复杂度O(n)，平均时间复杂度O(n^2)
func SelectSort(nums []int) {
	if len(nums) < 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

// MergeSort 稳不稳定看merge函数，最好是稳定的
// 非原地排序，所以没quicksort应用广泛
// 最好、最坏、平均时间复杂度均为O(nlogn)，空间复杂度为O(n)
// 从下而上
func merge(nums1, nums2 []int) []int {
	len1, len2 := len(nums1), len(nums2)
	i, j := 0, 0
	res := make([]int, 0, len1+len2)
	for i < len1 && j < len2 {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	if i < len1 {
		res = append(res, nums1[i:]...)
	} else {
		res = append(res, nums2[j:]...)
	}
	return res
}

func MergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	mid := len(nums) / 2
	leftPart := MergeSort(nums[:mid])
	rightPart := MergeSort(nums[mid:])
	return merge(leftPart, rightPart)
}

// QuickSort 不稳定的原地排序算法
// 从上而下
// 平均时间复杂度O(nlogn)，最坏时间复杂度O(n^2)
func QuickSort(nums []int, low, high int) {
	if low >= high {
		return
	}
	pivot := Partition(nums, low, high)
	QuickSort(nums, low, pivot-1)
	QuickSort(nums, pivot+1, high)
}

func Partition(nums []int, low, high int) int {
	pivot := nums[high]
	for low < high {
		for low < high && nums[low] <= pivot {
			low++
		}
		nums[high] = nums[low]
		for low < high && nums[high] >= pivot {
			high--
		}
		nums[low] = nums[high]
	}
	nums[low] = pivot
	return low
}

// 桶排序
// 条件：很容易划分成m个桶，桶跟桶之间有天然排序，数据在桶内分布均匀
// 常用于外部排序
// 每个桶包含一定范围内的数据
func BucketSort(nums []int) {
	k := len(nums) / 2
	buckets := make([][]int, k)
	for i := 0; i < k; i++ {
		buckets[i] = make([]int, 0)
	}
	for _, num := range nums {
		buckets[max(num/2-1, 0)] = append(buckets[max(num/2-1, 0)], num)
	}
	for _, bucket := range buckets {
		sort.Slice(bucket, func(i, j int) bool {
			return bucket[i] < bucket[j]
		})
	}
	index := 0
	for i := 0; i < k; i++ {
		for _, num := range buckets[i] {
			nums[index] = num
			index++
		}
	}
	return
}

// 计数排序
// 桶排序的特殊情况，每个桶内只包含一个值
func CountingSort(nums []int) {
	m := 0
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	counter := make([]int, m+1)
	for _, num := range nums {
		counter[num]++
	}
	for i := 0; i < m; i++ {
		counter[i+1] += counter[i]
	}

	n := len(nums)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		res[counter[nums[i]]-1] = nums[i]
		counter[nums[i]]--
	}
	copy(nums, res)
}

// 基数排序
// 要求数据能划分为高低位，位之间有递进关系，每一位的范围不能太大（因为依赖桶排序）

func main() {
	nums := []int{4, 1, 3, 2, 1, 2, 4, 2, 1, 2, 3, 1}
	CountingSort(nums)
	fmt.Println(nums)
}
