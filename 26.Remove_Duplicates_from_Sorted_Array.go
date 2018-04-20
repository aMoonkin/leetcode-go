// my solution
// [] is not considered firstly

func removeDuplicates(nums []int) int {
    if len(nums) == 0{
        return 0
    }
	p := 0 //position
	n := 1 //number of independent value
	lv := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] != lv{
			lv = nums[i]
			p++
			if i != p{
				nums[i], nums[p] = nums[p], nums[i]
			}
			n++

		}
	}
	return n
}
