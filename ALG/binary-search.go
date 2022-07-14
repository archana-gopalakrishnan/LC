func search(nums []int, target int) int {
    l := 0
    h := len(nums)-1
    
   // Have mid calc in the for loop instead of outside it to save time 
    
    for l<=h{
         m := (l+h)/2
        if target == nums[m]{
            return m
        } else if target < nums[m]{
            h = m - 1
        }else{
            l = m + 1
        }
    }
    return -1
}
