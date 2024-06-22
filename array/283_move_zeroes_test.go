package array

//https://leetcode.cn/problems/move-zeroes/solution/yi-dong-ling-by-leetcode-solution/
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
https://leetcode.cn/problems/move-zeroes/description/
review
refresher
把零移动到最后
[0, 1, 0, 3, 12]
*/
func MoveZeroes(nums []int) {
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}

/*
快慢双指针，慢指针停留在0位置，快指针发现非0位置时交换
*/
func moveZeroes(nums []int) {
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}

/*
滚雪球思路
https://leetcode.com/problems/move-zeroes/solutions/172432/the-easiest-but-unusual-snowball-java-solution-beats-100-o-n-clear-explanation/
简而言之就是只遍历一遍
雪球大小表示0数量
i左侧的i-snowballsize区域都为0
通俗而言就是保证雪球的头部为0，尾部i-snowballsize非0
*/
func snowBall(nums []int) {
	snowBallSize := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			snowBallSize++
		} else if snowBallSize > 0 {
			//头部为0 nums[i]=0
			//尾部为非0 nums[i-snowballsize]=nums[i]
			nums[i], nums[i-snowBallSize] = 0, nums[i]
		}
	}
}

func TestMove(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		wantRes []int
	}{
		{
			name:    "1",
			nums:    []int{1, 0, 1},
			wantRes: []int{1, 1, 0},
		},
		{
			name:    "2",
			nums:    []int{0},
			wantRes: []int{0},
		},
		{
			name:    "3",
			nums:    []int{0, 1, 0, 3, 12},
			wantRes: []int{1, 3, 12, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MoveZeroes(tt.nums)
			assert.Equal(t, tt.wantRes, tt.nums)
		})
	}
}
