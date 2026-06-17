package kata
​
import(
"strconv"
)
​
func PrinterError(s string) string {
  count:= 0
   for i:=0; i <len(s); i++{
     if s[i] < 'a' || s[i] > 'm'{
       count++
     }
   }
  return strconv.Itoa(count) + "/" + strconv.Itoa(len(s))
}
​