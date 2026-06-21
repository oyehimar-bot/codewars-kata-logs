package kata
​
import "strconv"
​
func BinToDec(bin string) int {
  // your code here
  b, _ := strconv.ParseInt(bin, 2, 64)
  return int(b)
}
​