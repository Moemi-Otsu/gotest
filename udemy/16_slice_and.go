// スライスのmakeとcap
// キャパシティは、あらかじめメモリを決めておく必要があるような開発に有効
package main

import "fmt"

func main() {
  // 3 = レングス, 5 = キャパシティ（メモリに確保しておくキャパシティ）
  n := make([]int, 3, 5)
  fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
  n = append(n, 0, 0)
  fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
  n = append(n, 1, 2, 3, 4, 5)
  fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

  // 引数が１つの場合、レングスもキャパシティもその引数になる
  a := make([]int, 3)
  fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)

  b := make([]int, 0)
  var c []int
  fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
  fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

//   c = make([]int, 5)
  c = make([]int, 0, 5)
  for i := 0; i < 5; i++{
    c = append(c, i)
    fmt.Println(c)
  }
  fmt.Println(c)
}