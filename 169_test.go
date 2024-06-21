package main

/*

假设当前数为众数, 遇到相同的数字则加一,否则减一
若前n个票和为0(互相抵消), 则设下一个数为当前数
重复1, 2 最后的当前数为众数
*/
func majorityElement(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	val, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			val = v
			count++
		} else if val == v {
			count++
		} else {
			count--
		}
	}
	return val
}
