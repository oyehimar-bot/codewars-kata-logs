package kata
​
import "strconv"
​
func BonusTime(salary int, bonus bool) string {
  // Your code here
  if bonus == true {
    salary *= 10
  }
  return "£" + strconv.Itoa(salary)
}