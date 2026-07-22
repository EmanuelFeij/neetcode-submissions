func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, i := range nums {
		counter[i] = counter[i] + 1
	}

	ret := make([]int, 0, k)
	i := 0

	for i < k {
		maxKey, maxValue := -1001, 0

		for k, v := range counter {
			if v > maxValue {
				maxKey = k
				maxValue = v
			}
		}

		delete(counter, maxKey)
		ret = append(ret, maxKey)
		i++
	}

	return ret
}
