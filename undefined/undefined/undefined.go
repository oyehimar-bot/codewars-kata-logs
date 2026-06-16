package kata
​
func Solution(word string) string {
  res:= ""
  for _, ch:= range word{
    res = string(ch) + res
  }
  return res
}