package main

import "fmt"

func main() {

  // 配列の書き方1
  var a [2]int
  a[0] = 100
  a[1] = 200
  fmt.Println(a)

  /*
  // 配列の書き方2
  // 配列の場合は、[2]int この部分が型になる
  var b [2]int = [2]int{100, 200}
  b = append(b, 300) //このように配列に追加しようとするとエラー
  fmt.Println(b)
  */

  // 配列はappendできない、スライス[] はappendできる。
  var b []int = []int{100, 200}
  b = append(b, 300)
  fmt.Println(b)
}