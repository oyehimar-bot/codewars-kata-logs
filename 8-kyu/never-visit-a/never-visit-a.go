83: "orange",
84: "grape",
85: "cherry",
86: "grape",
87: "cherry",
88: "pear",
89: "cherry",
90: "apple",
91: "kiwi",
92: "banana",
93: "kiwi",
94: "banana",
95: "melon",
96: "banana",
97: "melon",
98: "pineapple",
99: "apple",
100: "pineapple",
  }
  
  Sum:= func(nb int) int {
    sum := 0
  for nb > 0 {
    mod:= nb%10
    sum += mod
    nb /= 10
      }
    return sum
    }
  
  
 for {
   n = n - Sum(n)
   
   if fruit, ok := fruits[n]; ok {
     return fruit
   }
 }
}