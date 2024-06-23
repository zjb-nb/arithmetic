package array

//https://leetcode.cn/problems/merge-sorted-array/
import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
*
review/refresher
合并两个非严格递增数组，结果数组按照非严格递增顺序排列,nums2合并到nums1中
nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
[1,2,2,3,5,6]
*
*/
func Merge(nums1 []int, m int, nums2 []int, n int) {
	count, i, j := m+n-1, m-1, n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[count] = nums1[i]
			i--
		} else {
			nums1[count] = nums2[j]
			j--
		}
		count--
	}
}

/*
nums1的长度位m，nums2的长度为n
合并后的数组长度为,m+n
for 循环m+n次
移动双指针 num1[0]  nums[0]
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	var res = make([]int, 0, m+n)
	var i, j = 0, 0
	for i < m || j < n {
		if i < m {
			if j < n && nums1[i] > nums2[j] {
				res = append(res, nums2[j])
				j++
				continue
			}
			res = append(res, nums1[i])
			i++
			continue
		}

		if j < n {
			res = append(res, nums2[j])
			j++
		}
	}
	copy(nums1, res)
}

func merge2(nums1 []int, m int, nums2 []int, _ int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

/*
逆向双指针的前提是后面有足够的空空间
需要注意nums1被提前消耗完的情况，所以for条件必须是另一个
*/
func merge3(nums1 []int, m int, nums2 []int, n int) {
	var count, i, j = m + n - 1, m - 1, n - 1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[count] = nums1[i]
			i--
		} else {
			nums1[count] = nums2[j]
			j--
		}
		count--
	}
}

func TestMerge(t *testing.T) {
	var tests = []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		Want  []int
	}{
		{
			name:  "[1,2,2,3,5,6]",
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			Want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:  "[1]",
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			Want:  []int{1},
		},
		{
			name:  "[1]",
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			Want:  []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Merge(tt.nums1, tt.m, tt.nums2, tt.n)
			assert.Equal(t, tt.Want, tt.nums1)
		})
	}
}

/*
14.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
*/
