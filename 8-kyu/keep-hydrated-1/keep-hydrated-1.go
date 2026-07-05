package kata
​
import "math"
​
func Litres(time float64) int {
  //The fun begins here.
  return int(math.Floor(time * 0.5))
}