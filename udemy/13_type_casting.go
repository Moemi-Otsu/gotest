package main

import (
  "fmt"
  "strconv"
)

func main() {

  var x int = 1
  // float64に型変換（cast）
  xx := float64(x)
  // %T type
  // %v value
  // %f 指数なしの小数
  fmt.Printf("%T %v %f\n", xx, xx, xx)


  var y float64 = 1.2
  // integer型に型変換(cast)
  yy := int(y)
  // %T type
  // %v value
  // %d integer
  fmt.Printf("%T %v %d\n", yy, yy, yy)


  // Goでは、上記と同じようにstring型をint型に変換できない
  // var s string = "14"
  // z := int(s) // エラー cannot convert s (type string) to type int

  // string => int への型変換はこうやる（strconv）
  var s string = "14"
  i, _ := strconv.Atoi(s) // Atoi = ascii to integer
  fmt.Printf("%T %v\n", i, i)


  // 文字列とは、バイト配列のかたまり
  h := "Hello World"
  fmt.Println(h[0]) // 72
  fmt.Println(string(h[0])) // h
}