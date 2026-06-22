package kata
​
​
func SetAlarm(employed, vacation bool) bool {
    switch{
    case employed && vacation:
      return false
    case employed && !vacation:
      return true
    case !employed && vacation:
      return false
    case !employed && !vacation:
      return false 
    default:
  return false
  }
}