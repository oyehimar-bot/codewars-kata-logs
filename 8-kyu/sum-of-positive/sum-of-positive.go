package kata
​
func PositiveSum(numbers []int) int {
  sum := 0
  for _, number := range numbers{
    if number < 0 {
      continue
    }
    sum += number
  }
  return sum
}
​