package kata
​
func AddLetters(letters []rune) rune {
  // your code here
  res := 0
  for _, letter := range letters {
    res += int(letter) - 96
  }
  res %= 26
  if res == 0{
    return 'z'
  }
  return rune(res + 96)
}