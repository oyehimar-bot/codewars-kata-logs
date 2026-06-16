package kata
import(
"strings"
)
func ToAlternatingCase(str string) string {
  res:= ""
  for i:= 0; i< len(str); i++{
    if str[i] >= 'a' && str[i]<= 'z'{
    res += strings.ToUpper(string(str[i]))    
  } else if  str[i] >= 'A' && str[i]<= 'Z'{
    res += strings.ToLower(string(str[i]))
    } else {
      res+=string(str[i]) 
    }
   
  }
  return res 
}