package main

import "sort"

// 4 ms	3.1 MB
// 我们需找出两个数组的交集元素，同时应与两个数组中出现的次数一致。这样就导致了我们需要知道每个值出现的次数，所以映射关系就成了<元素,出现次数>。
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2,nums1)
	}
	m := map[int]int{}
	for _,value := range nums1 {
		m[value] ++
	}
	finalArray := []int{}
	for _,v := range nums2{
		if m[v] >0 {
			finalArray = append(finalArray,v)
			m[v] --
		}
	}
	return finalArray
}

// 4 ms	2.8 MB
// 对于两个已经排序好数组的题，我们可以很容易想到使用双指针的解法~
func intersectByTwoPoint(nums1 []int, nums2 []int) []int  {
	sort.Ints(nums1)
	sort.Ints(nums2)
	length1,length2 := len(nums1),len(nums2)
	index1,index2 := 0,0
	intersect := []int{}
	for index1 < length1 && index2 <length2 {
		if nums1[index1] < nums2[index2] {
			index1 ++
		} else if nums1[index1] > nums2[index2] {
			index2 ++
		} else {
			intersect = append(intersect,nums1[index1])
			index1 ++
			index2 ++
		}
	}
	return intersect

}

// 4 ms	2.7 MB
// 解答中我们并没有创建空白数组，因为遍历后的数组其实就没用了。我们可以将相等的元素放入用过的数组中，就为我们节省下了空间。
// 即我们往用过的数据写值的位置 是肯定被遍历过的
func intersectByTwoPointNoArray(nums1 []int, nums2 []int) []int {
	i, j, k := 0, 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if  nums1[i] < nums2[j] {
			i++
		} else {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}

// 序号350
func main()  {

	// 如果nums2的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中。
	//那么就无法高效地对nums2进行排序，因此推荐使用方法一而不是方法二.在方法一中，nums2只关系到查询操作，
	//因此每次读取nums2中的一部分数据，并进行处理即可。
}

