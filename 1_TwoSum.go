func twoSum(nums []int, target int) []int {
   var res []int
   tempMap := make(map[int]int) 
   for i := 0; i < len(nums); i++ {
        k := target - nums[i]
        _,exist := tempMap[k]
        if(exist) {
            return []int{tempMap[k],i}
        } else {
            tempMap[nums[i]] = i
        }
   } 
   return res
}
