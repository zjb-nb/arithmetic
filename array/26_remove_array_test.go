package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//https://leetcode.cn/problems/remove-duplicates-from-sorted-array/

/*
review refresher
nums = [0,0,1,1,1,2,2,3,3,4]
res ->  5, nums = [0,1,2,3,4]
*/
func RemoveDuplicates(nums []int) int {
	panic("")
}

/*
1.笨办法，就是额外创造一个进行遍历

	  res 指针遇到相同+1
		遍历结束条件为len(nums)
		最后copy

2.快慢双指针
*/
func removeDuplicates(nums []int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		wantRes []int
		wantNum int
	}{
		{
			name:    "1",
			nums:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			wantRes: []int{0, 1, 2, 3, 4},
			wantNum: 5,
		},
		{
			name:    "2",
			nums:    []int{1, 1, 2},
			wantRes: []int{1, 2},
			wantNum: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num := RemoveDuplicates(tt.nums)
			ok := assert.Equal(t, tt.wantNum, num)
			if ok {
				assert.Equal(t, tt.wantRes, tt.nums[:num])
			}
		})
	}
}
