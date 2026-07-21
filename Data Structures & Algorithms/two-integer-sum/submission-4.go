import "slices"

func twoSum(nums []int, target int) []int {
	numsIdx := make(map[int][]int)


	for idx, num := range nums{
		numsIdx[num] = append(numsIdx[num], idx)
	}
	

	for num, idxs := range numsIdx {
		missingNum := target - num
		if missingNum == num && len(idxs) == 2{
			return idxs
		}

		if  ns, ok :=numsIdx[missingNum]; ok && missingNum != num {
			 ss := append([]int{idxs[0]}, ns[0])
			 slices.Sort(ss)
			 return ss
		}
	}
    
	return []int{}
}
