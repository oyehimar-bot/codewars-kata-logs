package kata
​
func LeastLarger(a []int, i int) int {
  target := a[i]
  resIndex:= -1
  smallestLarge:= int(^uint(0) >> 1)
  
  for idx, val := range a {
    if val > target && val < smallestLarge {
      smallestLarge = val
       resIndex = idx
    }
  }
  return resIndex
}
​
​
​
​
​
​
​
​
​
​
​
​
​
//  num := a[i]
  
// max := func(xx []int) int {
//   m:= a[0]
//   for _, ma:= range xx {
//     if ma > m{
//       m = ma