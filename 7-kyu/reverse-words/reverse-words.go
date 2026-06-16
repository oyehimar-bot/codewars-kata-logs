package kata
import (
"strings"
) 
func ReverseWords(str string) string {
  st := strings.Split(str, " ")
  for i, ch := range st{
    res:= ""
    for _, s:= range ch{
      res = string(s) + res
    }
    st[i] = res
  }
 return strings.Join(st, " ")
}