// Kadane's algorithm
// maxSum is first value of array (since output sum is sum of at least one array element)
// current Sum (curSum) is set to 0 since the array can have both positive and negative integers
// if the curSum value goes below zero, we can ignore the array elements before it
// explanantion video:

func maxSubArray(nums []int) int {
    maxSum := nums[0]
    curSum := 0
    
    for _, val := range(nums){
        if curSum < 0 {
            curSum = 0
        }
        curSum = curSum + val
    
        maxSum = Max(maxSum, curSum)
    }   
    return maxSum
}

func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

// Better implementation - had a faster run time but the first one required lesser memory
// Based on Nick White's video
// https://www.youtube.com/watch?v=jnoVtCKECmQ

func maxSubArray(nums []int) int {
  maxSum := nums[0]
  curSum := maxSum
  
  //loop from second element
  for j:=1; j<len(nums); j++ {
    // if adding next element makes curSum bigger than the next element, then assign the value to curSum
    // else discard curSum and directly make the value equal to the next element itself.
    curSum = Max(nums[j], (curSum + nums[j]))
    // if curSum is bigger then maxSum then assign that value to maxSum
    maxSum = Max(curSum, maxSum)
  }
  return maxSum
}
