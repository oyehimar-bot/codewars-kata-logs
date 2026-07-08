package kata
​
import "math"
​
func WallPaper(l, w, h float64) string {
  digitInWord := []string{
    "zero", "one", "two", "three", "four", "five",
    "six", "seven", "eight", "nine", "ten",
    "eleven", "twelve", "thirteen", "fourteen",
    "fifteen", "sixteen", "seventeen", "eighteen",
    "nineteen", "twenty",
  }
​
  if l == 0 || w == 0 || h == 0 {
    return "zero"
  }
​
  const rollArea = 5.2 
  
  wallArea := 2 * h * (l + w)
   
  rolls := int(math.Ceil((wallArea * 1.15) / rollArea))
​
  return digitInWord[rolls]
}
​