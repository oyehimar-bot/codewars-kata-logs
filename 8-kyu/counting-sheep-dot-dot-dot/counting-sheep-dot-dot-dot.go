package kata
​
func CountSheeps(numbers []bool) int {
  // your code here
  count:=0
  for _, number := range numbers {
    if number == true {
      count++
    }
  }
  return count 
}
​