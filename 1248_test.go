package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
review
给你一个整数数组 nums 和一个整数 k。如果某个连续子数组中恰好有 k 个奇数数字，
我们就认为这个子数组是「优美子数组」。
请返回这个数组中 「优美子数组」 的数目。

输入：nums = [1,1,2,1,1], k = 3
输出：2
解释：包含 3 个奇数的子数组是 [1,1,2,1] 和 [1,2,1,1] 。

输入：nums = [2,4,6], k = 1
输出：0
解释：数列中不包含任何奇数，所以不存在优美子数组。
*/
func NumberOfSubarrays(nums []int, k int) int {
	//TODO
	return 0
}

/*
1.滑动窗口 暴力统计
当窗口内的奇数个数=k时
left左侧的连续偶数段的任意一个，都可以作为子段，即优美数组的起点，
还要包括第一个奇数lefteventcnt+1
同理窗口右侧的连续偶数段，总子数组个数为(lefteventcnt+1)*(righteventcnt+1)
*/
func numberOfSubarrays(nums []int, k int) int {
	res, cout, left, right := 0, 0, 0, 0
	for right < len(nums) {
		if nums[right]%2 == 1 {
			cout++
		}
		right++
		if cout == k {
			temp := right
			//向右滑动，直到遇到下一个奇数的其实边界
			for right < len(nums) && nums[right]%2 == 0 {
				right++
			}
			//第k个奇数右边的偶数个数
			rightEventCnt := right - temp
			//第一个奇数左边偶数的个数
			leftEventCnt := 0
			//第一个奇数左边的leftEventCnt个偶数都可以作为优美数组的起点
			//第k个奇数右边rightEventCnt个偶数都可以作为优美数组的终点
			//所以该段一共有(lefteventcnt+1)*(rightEventCnt+1)个组合
			for nums[left]%2 == 0 {
				leftEventCnt++
				left++
			}
			res += (leftEventCnt + 1) * (rightEventCnt + 1)
			//此时left的指向第一个奇数
			left++
			cout--
		}
	}
	return res
}

/*
2.前缀和算法
先进行预处理奇数为1，偶数为0
求有几个子数组有k个奇数问题就变为了 ，有多少个[j,i]区间和为k的问题
s[j]-s[i] = k
[1,1,2,1,1] => [1,1,0,1,1]
s[i]为前i个的和，即
s[i]=> [1,2,2,3,4]
因为s[i] = s[i-1]+nums[i],所以需要 len(s) = len(nums)+1,即s[0]初始化为0
s[i]=>[0,1,2,2,3,4] k=>3
s[4]-s[0] = 3
s[5]-s[1]=3
所以有两个优美数组，再完善一下条件
有多少个满足 s[j]-s[i] = k =>

	s[i] = s[j] - k

这不就是求两数之和的公式吗，只不过求解换成了记录的哈希表存在几个 s[j]-k的值

对于 s[5] （以nums[4]为边界） 存在1个 边界s[1]满足 3=3-k =>[1,5]	---[1,0,1]
对于 s[4] （以nums[3]为边界） 存在1个 边界s[0]满足 3=3-k =>[0,4]	---[1,1,0,1]

问题就变成了如何查找 差的个数 ，用cout保存，如果 count [ s[5]-k ] 命中 把值累加
cout[0]=0 表示s[n]和为0的存在个数为0
*/
func numberOfSubarrays1(nums []int, k int) int {
	res := 0
	s := make([]int, len(nums)+1)
	cout := make([]int, len(s))

	for i, v := range nums {
		s[i+1] = s[i] + v%2
	}

	for _, v := range s {
		if v-k >= 0 {
			res += cout[v-k]
		}
		cout[v]++
	}
	return res
}

func TestSubarrays(t *testing.T) {
	tests := []struct {
		name    string
		data    []int
		k       int
		wantRes int
	}{
		{
			name:    "1",
			data:    []int{1, 1, 2, 1, 1},
			k:       3,
			wantRes: 2,
		},
		{
			name:    "2",
			data:    []int{2, 4, 6},
			k:       1,
			wantRes: 0,
		},
		{
			name:    "3",
			data:    []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2},
			k:       2,
			wantRes: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := NumberOfSubarrays(tt.data, tt.k)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
