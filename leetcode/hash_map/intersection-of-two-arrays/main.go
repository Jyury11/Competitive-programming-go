package main

func intersection(nums1 []int, nums2 []int) []int {
	hashMap := make(map[int]struct{}, len(nums1))
	ret := make([]int, 0, len(nums1))
	for _, num := range nums1 {
		if _, exists := hashMap[num]; !exists {
			hashMap[num] = struct{}{}
		}
	}

	okMap := make(map[int]struct{}, len(nums1))
	for _, num := range nums2 {
		if _, exists := hashMap[num]; exists {
			if _, ok := okMap[num]; !ok {
				ret = append(ret, num)
				okMap[num] = struct{}{}
			}

		}
	}

	return ret
}
