package main

import "fmt"

func main() {

// %T = 型
// %v = 値
// %t = 単語、true または false 指定した型じゃないと正しく表示されない

//   var t, f bool = true, false
  t, f := true, false
  fmt.Printf("%T %v %t\n", t, t, t)
  fmt.Printf("%T %v %t\n", f, f, f)

// 論理演算子 && （and）
  fmt.Println(true && true) // 真かつ真 = 真
  fmt.Println(true && false) // 真かつ偽 = 偽
  fmt.Println(false && false) // 偽かつ偽 = 偽

// 論理演算子 || （or）
  fmt.Println(true || true) // 真もしくは真 = 真
  fmt.Println(true || false) // 真もしくは偽 = 真
  fmt.Println(false || false) // 偽もしくは偽 = 偽

// 論理演算子 ! （not）
  fmt.Println(!true) // false 条件の反対の結果
  fmt.Println(!false) // true 条件の反対の結果
}