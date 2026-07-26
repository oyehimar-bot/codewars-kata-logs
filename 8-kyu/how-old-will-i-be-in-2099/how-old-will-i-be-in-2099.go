package kata
​
import "fmt"
​
func calculateAge(born, year int) string {
  age := year - born
  if age > 1 {
    return fmt.Sprintf("You are %d years old.", age)
  } else if age < -1 {
    return fmt.Sprintf("You will be born in %d years.", -age)
  } else if age == 1 {
    return "You are 1 year old."
  } else if age == -1 {
    return "You will be born in 1 year."
  }
  return "You were born this very year!"// enter your code here.
}
​