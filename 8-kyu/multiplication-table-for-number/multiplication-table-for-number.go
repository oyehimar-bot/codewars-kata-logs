package kata
​
import "strconv"
func MultiTable(number int) string {
  //good luck
  result := ""
  for i := 1; i <= 10; i++ {
    line := strconv.Itoa(i) + " * " + strconv.Itoa(number) + " = " + strconv.Itoa(i*number)
    if i < 10 {
      line += "\n"
    }
    result += line
  }
  return result
}