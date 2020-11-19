// mapとは 連想配列
package main

import "fmt"

func main() {
  m := map[string]int{"apple": 100, "banana": 200}
  fmt.Println(m)

  // キーを表示することで 100 が取り出せる。
  fmt.Println(m["apple"])

  // banana の内容を 300 で上書き
  m["banana"] = 300
  fmt.Println(m)

  // 追加する
  m["new"] = 500
  fmt.Println(m)

  // mapに存在しないキーを指定すると 0 が返る
  fmt.Println(m["nothing"])

  // 100 true 返り値がある場合はこのように２つ書ける??ここよくわからん
  v, ok := m["apple"]
  fmt.Println(v, ok)

  // 0 false
  v2, ok2 := m["nothing"]
  fmt.Println(v2, ok2)

  // m2["pc"] = 5000
  m2 := make(map[string]int)
  m2["pc"] = 5000
  fmt.Println(m2)

  // panic: assignment to entry in nil map
  // 宣言はしているものの、メモリ上にmapがないのでエラーに
  /*
  var m3 map[string]int
  m3["pc"] = 5000
  fmt.Println(m3)
  */

  // var宣言した場合は、
  // sliceのときでも、mapのときでもnilが返る。
  var s []int
  if s == nil{
    fmt.Println("Nil")
  }
}