package main

func main() {
	array := []int{1, 2, 3}
	permute(array)
}

func permute(nums []int) [][]int {

	// length := len(nums)
	// count := 1
	// for length > 0 {
	// 	count *= length
	// 	length--
	// }

	// arr := make([][]int, count)
	// for i := 0; i < count; i++ {
	// 	arr[i] = make([]int, length)
	// }

	used := make([]bool, len(nums))
	p := []int{}
	res := [][]int{}
	generatePermutation(nums, 0, p, &res, &used)
	return nil
}

func generatePermutation(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			p = append(p, nums[i])
			generatePermutation(nums, index+1, p, res, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
}
