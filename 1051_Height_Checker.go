func heightChecker(heights []int) int {
   copied := make([]int,len(heights))
   copy(copied,heights)
   sort.Ints(copied)
   count := 0
   for i := 0; i < len(heights); i++ {
      if heights[i] != copied[i] {
        count++
      }
   }

   return count  
}
