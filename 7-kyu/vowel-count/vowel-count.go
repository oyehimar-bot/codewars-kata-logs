package kata
​
func GetCount(str string) (count int) {
  for _, s:= range str {
    if s == 'a' || s =='e'|| s == 'i' || s == 'o' || s == 'u'{
      count++
    }
  }
  return count
}