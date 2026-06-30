package kata
​
func CalculateYears(years int) (result [3]int) {
  // Write your solution here
  catYears, dogYears := 0, 0
  if years == 1 {
    catYears = 15
    dogYears = 15
  } else if years == 2 {
    catYears = 15 + 9
    dogYears = 15 + 9
  } else {
    catYears = 15 + 9 + (years - 2) * 4
    dogYears = 15 + 9 + (years - 2) * 5
  }
  result[0], result[1], result[2] = years, catYears, dogYears
  return 
}