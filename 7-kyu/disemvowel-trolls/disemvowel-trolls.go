package kata
​
func Disemvowel(comment string) string {
    res := []rune{} // use rune slice for efficient appending
​
    for _, s := range comment {
        switch s {
        case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
            // skip vowels
            continue
        default:
            res = append(res, s)
        }
    }
​
    return string(res)
}
​