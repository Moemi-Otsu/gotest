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
}