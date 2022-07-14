
func searchInsert(nums []int, target int) int {
    // implement binary search
    // return the mid value
    
    l := 0
    h := len(nums)
    
    for l < h{
        m := l + (h-l)/2
        if target <= nums[m]{
            h = m
        } else{
            l = m+1
        }
    }
    
    return h
    
}

 
// alternate solution

func searchInsert(nums []int, target int) int {
    // implement binary search
    // return the mid value
    
    l := 0
    h := len(nums)
    
    for l < h{
        m := l + (h-l)/2
        if target == nums[m]{
            return m
        } else if target < nums[m]{
            h = m
        } else{
            if (m == len(nums)-1) || (target <= nums[m+1]) { // lines 38 - 40 were make/break
				    return m + 1
		      	}
            l = m+1
        }
    }
    
    return l
    
}
