package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// review
/*
https://leetcode.cn/problems/rotate-array/description/
nums = [1,2,3,4,5,6,7], k = 3
[5,6,7,1,2,3,4]
*/
func Rotate(nums []int, k int) {
	count, n := 0, len(nums)
	k = k % n
	for i := 0; i < n; i++ {
		temp, next := nums[i], i
		for {
			next = (next + k) % n
			temp, nums[next] = nums[next], temp
			count++
			if next == i {
				break
			}
		}
		if count == n {
			return
		}
	}
}

/*
nums = [1,2,3,4,5,6,7], k = 3
[5,6,7,1,2,3,4]

先整体旋转，再部分旋转
[7,6,5 ]   [4,3,2,1]
*/
func rotate(nums []int, k int) {
	//这一步为了解决k>len(nums)
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

// 二分翻转数组
func reverse(a []int) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

/*
环状数组解法
交换成功后temp需要存放当前要跳到位置的值，用于下一次交换1-》 3  temp=nums[1]  temp,nums[1+3] = nums[1+3],temp
k为奇数

	一次遍历即可，结束条件为 p==0 或者cnt==len(nums)

k为偶数

	第一次遍历到p==0条件时，发现偶数位置都被跳过，需要外面再套一层循环，结束条件为cnt==len(nums)

综上，二层循环，第一层确保k为偶数能全部遍历，第二层循环执行交换主逻辑
*/
func rotate1(nums []int, k int) {
	n := len(nums)
	k %= n
	cnt, p := 0, 0

	for i := 0; i < n; i++ {
		temp := nums[i]
		p = i
		for {
			p = (p + k) % n
			nums[p], temp = temp, nums[p]
			cnt++
			if p == i {
				break
			}
		}
		if cnt == n {
			break
		}
	}

}
func TestRotate(t *testing.T) {
	tests := []struct {
		name    string
		data    []int
		k       int
		wantRes []int
	}{
		{
			name:    "1",
			data:    []int{1, 2, 3, 4, 5, 6, 7},
			k:       3,
			wantRes: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:    "2",
			data:    []int{-1, -100, 3, 99},
			k:       2,
			wantRes: []int{3, 99, -1, -100},
		},
		{
			name:    "2",
			data:    []int{1, 2},
			k:       1,
			wantRes: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rotate(tt.data, tt.k)
			assert.Equal(t, tt.wantRes, tt.data)
		})
	}
}
