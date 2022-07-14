/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

/*
//got run time error: out of memory
func firstBadVersion(n int) int {
    if isBadVersion(n) && !isBadVersion(n-1) {
        return n
    } else{
        return firstBadVersion(n-1)
    }
}
*/

// use binary search instead

func firstBadVersion(n int) int {
   l := 0
    h := n 
// time limit execeeded when using h := n-1 and for l<= h
   
   for l < h {
        m := l + (h-l)/2
        if !isBadVersion(m){
            l = m + 1
        }else{
            h = m
        }
   }
   return l
}

