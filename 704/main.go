package main

func main() {
	array := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	search(array, target)
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, mid, high := 0, (len(nums)-1)/2, len(nums)-1
	for low <= high {
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		}
		mid = (low + high) / 2
	}
	return -1
}
