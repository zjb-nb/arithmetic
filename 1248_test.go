package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
2种方法
review https://leetcode.cn/problems/count-number-of-nice-subarrays/description/
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
	/*
			1.全部转为0、1 求k个奇数就转为范围和s[i,j]为k
			2. s[j]=s[j-1]+nums[j] 因此为防止越界s[0]=0 s[0]表示以nums[0]为边界(不包括)的范围和
		  3. s[j]-s[i] = k   ==> s[i]= s[j]-k 即轮询s到s[j]时，是否存在s[i]另s[j]-k=s[i]
			有几个s[i]就有几个子数组（以nums[j]为边界），这不就是两数之和(s[j]和-k)的问题？
			我们用频次数组cnt保存s[i]   如果s[j]-v>=0(因为s[i]>=0) 时，如果 cnt[ s[j]-v ] 存在就res+=xxxx
			s[0]=0
			for k,v := range nums {
			  s[k+1]=xxxx
			}
		  cnt  map[][]
			for _,v:= range s {
			  if v-k>=0 {
				  res += cnt[v-k]
				}
			  cnt[v]++
			}
	*/
	res := 0
	s := make([]int, len(nums)+1)
	for i, v := range nums {
		s[i+1] = s[i] + v%2
	}
	cnt := make(map[int]int)
	for _, v := range s {
		if v-k >= 0 {
			res += cnt[v-k]
		}
		cnt[v]++
	}
	return res
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

	将奇数偶数看成1、0
	k=3
	[1,1,2,1,1] = [1,1,0,1,1]
	有k个奇数问题=》 有多少个区间[i,j]和为3
	s为以i为边界（不包括i）的和
	s[j] = s[j-1]+nums[j]
	s = [0,1,2,4,5,6]
	区间和 就是 s[j]-s[i] =k (两数之和问题 s[j]/-s[i])
	s[i] = s[j]-k
	我们考虑以j为边界的优美子数组时，只要统计有多少个s[i]
	因此建立频次数组cnt记录s[j]的个数 那么 轮询s时 cnt[s[i]-k]即可得到
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
