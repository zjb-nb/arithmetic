package array

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// review
func threeSum1(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for cur := 0; cur < len(nums)-2; cur++ {
		if nums[cur] > 0 {
			return res
		}
		if cur > 0 && nums[cur] == nums[cur-1] {
			continue
		}

		l, r := cur+1, len(nums)-1
		for l < r {
			if nums[cur]+nums[l]+nums[r] == 0 {
				res = append(res, []int{nums[cur], nums[l], nums[r]})
				for l < r && nums[l+1] == nums[l] {
					l++
				}
				l++
				for l < r && nums[r-1] == nums[r] {
					r--
				}
				r--
			} else if nums[cur]+nums[l]+nums[r] < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

//	 寻找三数之和为0 ，且去重，返回组合，不是下标
//		-4  -1 -1 0 1 2
func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for current := 0; current < len(nums)-2; current++ {
		// 从小到大排序，第一个大于0，总和肯定不为0
		if nums[current] > 0 {
			return res
		}
		if current > 0 && nums[current] == nums[current-1] {
			continue
		}
		//进行求和,i,j双指针
		left, right := current+1, len(nums)-1
		for left < right {
			if nums[current]+nums[left]+nums[right] == 0 {
				//跳过重复
				res = append(res, []int{nums[current], nums[left], nums[right]})
				for nums[left] == nums[left+1] && left < right {
					left++
				}
				for nums[right] == nums[right-1] && right > left {
					right--
				}
				left++
				right--
			} else if nums[current]+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}
		}

	}
	return res
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)

	res := [][]int{}

	for k, v := range nums {
		if nums[k] > 0 {
			return res
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		L := k + 1
		R := len(nums) - 1
		for L < R {
			if v+nums[L]+nums[R] == 0 {
				res = append(res, []int{v, nums[L], nums[R]})
				for (L < R) && nums[L] == nums[L+1] {
					L++
				}
				for (L < R) && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if v+nums[L]+nums[R] > 0 {
				R--
			} else {
				L++
			}
		}
	}
	return res
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name    string
		data    []int
		wantRes [][]int
	}{
		{
			name: "a",
			data: []int{-1, 0, 1, 2, -1, -4},
			wantRes: [][]int{
				[]int{-1, -1, 2},
				[]int{-1, 0, 1},
			},
		},
		{
			name:    "b",
			data:    []int{0, 1, 1},
			wantRes: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := threeSum1(tt.data)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
