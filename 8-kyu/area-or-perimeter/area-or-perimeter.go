package kata
​
func AreaOrPerimeter(l, w int) int {
  if l == w {
    return l * w
  }
  return 2 * (l + w) // Enter solution here
}