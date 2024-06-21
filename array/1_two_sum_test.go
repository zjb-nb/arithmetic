package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	input  []int
	target int
}

//https://leetcode.cn/problems/two-sum/

// review 给一个数字求两数之和，返回两下标
func TwoSum(nums []int, target int) []int {
	panic("")
}

/*
构造map -> map[ target-本身 ]目标下标
存在则直接返回
不存在则放入map
空间换时间
O1 复杂度
*/
func twoSum(nums []int, target int) []int {
	tmp_map := make(map[int]int)
	for k, v := range nums {
		if exist_target, val := tmp_map[target-v]; val {
			return []int{k, exist_target}
		}
		tmp_map[v] = k
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		val     TestData
		wantRes []int
	}{
		{
			name: "first",
			val: TestData{
				input:  []int{2, 7, 11, 15},
				target: 9,
			},
			wantRes: []int{1, 0},
		},
		{
			name: "second",
			val: TestData{
				input:  []int{3, 2, 4},
				target: 6,
			},
			wantRes: []int{2, 1},
		},
		{
			name: "third",
			val: TestData{
				input:  []int{3, 3},
				target: 6,
			},
			wantRes: []int{1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := TwoSum(tt.val.input, tt.val.target)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
