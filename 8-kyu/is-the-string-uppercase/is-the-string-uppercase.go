package kata
​
type MyString string
​
func (s MyString) IsUpperCase() bool {
  // Your code here!
  for i := 0; i < len(s); i++{
      if rune(s[i]) >= 'a'{
      return false
    }
  }
  return true
}