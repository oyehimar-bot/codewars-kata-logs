package kata
​
​
func SequenceSum(start, end, step int) int {
  num := ((end - start)/step) + 1
  sum := (num * (start + end))/2
  if end < start{
    return 0
    } else if (end - start) % step != 0 {
        end1 := start + step * ((end-start)/step)
          return (num * (start + end1))/2
  }
  return sum
}
​