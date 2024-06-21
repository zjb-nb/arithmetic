package main

import "testing"

/*
这里有 n 个航班，它们分别从 1 到 n 进行编号。
有一份航班预订表 bookings ，表中第 i 条预订记录
bookings[i] = [firsti, lasti, seatsi]
意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上
预订了 seatsi 个座位。
请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。

输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
输出：[10,55,45,25,25]
解释：
航班编号        1   2   3   4   5
预订记录 1 ：   10  10
预订记录 2 ：       20  20
预订记录 3 ：       25  25  25  25
总座位数：      10  55  45  25  25
因此，answer = [10,55,45,25,25]
*/

/*
1.暴力求解
2.利用差分数组，差分数组的前缀和就是原数组
*/
func corpFlightBookings(bookings [][]int, n int) []int {
	flights := make([]int, n)
	for _, v := range bookings {
		if v[0] > 0 {
			flights[v[0]-1] += v[2]
		}
		if v[1] < n {
			flights[v[1]] -= v[2]
		}
	}
	res := make([]int, n+1)
	for k, v := range flights {
		res[k+1] = res[k] + v
	}
	return res[1:]
}

func TestFlight(t *testing.T) {
	t.Log(corpFlightBookings([][]int{
		[]int{1, 2, 10},
		[]int{2, 3, 20},
		[]int{2, 5, 25},
	}, 5))
}
