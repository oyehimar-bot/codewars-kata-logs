package kata
​
func GetSum(a, b int) int {
  if b>a{
    return (a+b)*(b-a+1)/2
  }
  return (a+b)*(-b+a+1)/2
}