package kata
​
import "math"
​
func CornerCircle(n int) float32 {
  res := float64(n) * (3 - 2*math.Sqrt(2))
  return float32(math.Round(res*100) / 100)
}
​