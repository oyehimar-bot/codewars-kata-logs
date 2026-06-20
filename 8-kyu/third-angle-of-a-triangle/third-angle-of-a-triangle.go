package kata
​
func OtherAngle(a int, b int) int {
  // your code here
  c := 180 - (a + b)
  if c < 0{
    c = -c
  }
  return c
}