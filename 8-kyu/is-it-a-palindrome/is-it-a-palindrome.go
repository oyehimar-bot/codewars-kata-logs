package kata
​
import "strings"
​
func IsPalindrome(str string) bool {
  str = strings.ToLower(str)
  res := ""
  for _, p := range str {
    res = string(p) + res
  }
  return res == str
}