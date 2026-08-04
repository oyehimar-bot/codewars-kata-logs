package kata
​
func Enough (cap, on, wait int) int {
  // Your solution here
  if (on + wait - cap) < 0{
    return 0
  }
  return on + wait - cap
}