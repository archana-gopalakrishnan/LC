func containsDuplicate(nums []int) bool {
    
    for i, val := range(nums){
        count := 0
        for j:=i; j < len(nums); j++ {
            if nums[j] == val{
                count++
            }
            if count >= 2{
                return true
            }
        }
    }
   return false
    
}
