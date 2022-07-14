func containsDuplicate(nums []int) bool {
    count := 0
    
    for _, val := range(nums){

        for j:=0; j < len(nums); j++ {
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
