package kata
​
func Factorial(n int) int {
  // put your code here
  if n <= 1{
    return 1
  }
  res := 1
  for i:= 1; i <= n; i++{
    res *= i
  }
  return res
}