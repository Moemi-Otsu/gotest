package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println("Hello World")
  fmt.Println("Hello " + "World")
  fmt.Println("Hello World"[0]) // 72
  fmt.Println(string("Hello World"[0]))

  var s string = "Hello World"
  s = (strings.Replace(s, "H", "X", 1))
  fmt.Println(s)
  fmt.Println(strings.Contains(s, "World"))

  // バッククオートで囲うと、エディタの改行がそのまま反映できる
  fmt.Println(`Test
Test`)

  // ダブルクオートの中にダブルクオート描きたい場合は
  // \（バックスラッシュ）でエスケープ
  fmt.Println("\"")

  // バックスラッシュめんどいときは、バッククオートでエスケープ
  fmt.Println(`"`)
}