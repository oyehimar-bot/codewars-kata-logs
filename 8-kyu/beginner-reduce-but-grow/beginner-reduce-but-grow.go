package kata
​
func Grow(arr []int) int{
  a:= 1
  for _, ar := range arr {
    a *= ar
  }
  return a
}