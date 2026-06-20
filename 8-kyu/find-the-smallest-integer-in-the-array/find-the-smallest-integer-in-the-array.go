package kata
​
func SmallestIntegerFinder(numbers []int) int {
  min := numbers[0]
  for _, number := range numbers{
    if number < min {
      min = number
    }
  }
  return min
}