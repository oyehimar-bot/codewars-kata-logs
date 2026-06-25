package kata
​
​
func Invert(arr []int) []int {
  a:= make([]int, len(arr))
  
  for i, v:= range arr {
    a[i] = -v
  }
  return a
}
​