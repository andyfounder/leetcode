package main

func main() {
	array1 := []int{1, 2, 3, 0, 0, 0}
	array2 := []int{2, 5, 6}
	num1 := 3
	num2 := 3
	merge(array1, num1, array2, num2)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// slice := make([]int, 0)
	// if n != 0 {
	// 	i := 0
	// 	j := 0
	// 	for i < m && j < n {
	// 		if nums1[i] < nums2[j] {
	// 			slice = append(slice, nums1[i])
	// 			i++
	// 		} else {
	// 			slice = append(slice, nums2[j])
	// 			j++
	// 		}
	// 	}
	// 	if i == m {
	// 		for j < n {
	// 			slice = append(slice, nums2[j])
	// 			j++
	// 		}
	// 	}
	// 	if j == m {
	// 		for j < m {
	// 			slice = append(slice, nums1[i])
	// 			i++
	// 		}
	// 	}
	// 	i = 0
	// 	for i, s := range slice {
	// 		nums1[i] = s
	// 		i++
	// 	}
	// }

	// slice := make([]int, m+n)
	// i := 0
	// j := 0
	// k := 0
	// for i < len(nums1) && nums1[i] != 0 {
	// 	if j < n && nums1[i] > nums2[j] {
	// 		slice[k] = nums2[j]
	// 		j++
	// 	} else {
	// 		slice[k] = nums1[i]
	// 		i++
	// 	}
	// 	k++
	// }
	// for j < n {
	// 	slice[k] = nums2[j]
	// 	j++
	// 	k++
	// }
	// for i, s := range nums1 {
	// 	nums1[i] = s
	// }
	for i := len(nums1) - 1; i >= 0; i-- {
		if m > 0 && n > 0 {
			if nums1[m-1] < nums2[n-1] {
				nums1[i] = nums2[n-1]
				n--
			} else {
				nums1[i] = nums1[m-1]
				m--
			}
		} else if n > 0 {
			nums1[i] = nums2[n-1]
			n--
		}
	}
}
