package main

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	containsNearbyDuplicate(nums, k)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 || len(nums) == 1 {
		return false
	}
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, flag := m[nums[i]]; flag {
			return true
		}
		m[nums[i]] = true
		if len(m) == k+1 {
			delete(m, nums[i-k])
		}
	}
	return false
}
