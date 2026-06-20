package kata
​
import(
  "sort"
  "strings"
)
func TwoSort(arr []string) string {
  sort.Strings(arr)
  b:= []string{}
  for i:= 0; i< len(arr[0]); i++{
    b= append(b, string(arr[0][i]))
  }
  return strings.Join(b, "***")
}