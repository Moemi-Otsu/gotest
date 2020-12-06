package main

import (
   "fmt"
   "sort"
)

func main() {
  // Q1 . 以下のスライスから一番小さい数を探して出力するコードを書いてください。
  l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
  sort.Sort(sort.IntSlice(l))
  fmt.Println(l[0])

  // Q2. 以下の果物の価格の合計を出力するコードを書いてください。
  m := map[string]int{
      "apple":  200,
      "banana": 300,
      "grapes": 150,
      "orange": 80,
      "papaya": 500,
      "kiwi":   90,
  }
  sum := 0
  for _, v := range m{
    sum += v
  }
  // fmt.Printf("%T %v", m, m)
  fmt.Println(sum)
}