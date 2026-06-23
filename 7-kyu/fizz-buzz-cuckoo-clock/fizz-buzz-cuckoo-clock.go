package kata
​
import (
  "strconv"
)
​
func FizzBuzzCuckooClock(time string) string {
  // your code here
  hr, _:= strconv.Atoi(time[:2])
  min, _:= strconv.Atoi(time[3:])
  displayhr := hr % 12
  if displayhr == 0 {
    displayhr = 12
  }
  if min == 0 {
    res := "Cuckoo"
    for i := 1; i < displayhr; i++ {
      res += " Cuckoo"
    }
    return res
  }
  if min == 30 {
    return "Cuckoo"
  }
  
  if min%3 == 0 && min%5 == 0 {
    return "Fizz Buzz"
  } else if min%3 == 0 {
    return "Fizz"
  } else if min%5 == 0 {
    return "Buzz"
  }
  return "tick"
}
​