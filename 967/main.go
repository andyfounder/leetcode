package main

func main() {
	n := 3
	k := 7
	numsSameConsecDiff(n, k)
}

func numsSameConsecDiff(n int, k int) []int {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if n == 1 {
		return []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}
	n--
	i := 0
	for n > 0 {
		res := []int{}
		for i < len(nums) {
			temp1 := nums[i]%10 + k
			temp2 := nums[i]%10 - k
			if temp1 < 10 {
				res = append(res, nums[i]*10+temp1)
			}
			if temp2 >= 0 {
				res = append(res, nums[i]*10+temp2)
			}
			i++
		}
		nums, i = res, 0
		n--
	}
	return nums
}
