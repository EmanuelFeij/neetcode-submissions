

func twoSum(nums []int, target int) []int {
	numsIdx := make(map[int]int)

for idx, num := range nums{
	numsIdx[num] = idx
}

	for idx, num := range nums {
		missingNum := target - num
		if  ns, ok :=numsIdx[missingNum]; ok && idx != ns {
			 return []int{ idx, ns}
		}
		
	}
	
    	return []int{}
}
