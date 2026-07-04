package kata
​
import "strings"
​
func ReverseWords(str string) string {
  st := strings.Split(str, " ")
  for left, right := 0, len(st)-1; left < right; left, right = left+1, right-1 {
        st[left], st[right] = st[right], st[left]
    }
  return strings.Join(st, " ") // reverse those words
}