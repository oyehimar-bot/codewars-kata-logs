package kata
​
import "strings"
​
func toWeirdCase(str string) string {
  // Your code here and happy coding!
  str = strings.ToLower(str)
  toSlice := strings.Fields(str)
  res := ""
  for i :=0; i < len(toSlice);i++{
    word := toSlice[i]
    for j :=0; j < len(word);j++{
      if j & 1 == 0 {
        res += strings.ToUpper(string(word[j]))
      } else {
        res += string(word[j])
      }
      if j == len(word)-1{
        res += " "
      }
    }
  }
  return res[:len(res)-1]
}