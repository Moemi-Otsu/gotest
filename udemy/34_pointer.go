package main

import "fmt"

func one(x *int){
  *x = 1
}

func main() {
  // デリファレンス
  var n int = 100
  one(&n)
  fmt.Println(n)
  fmt.Println(*&n)

//   // メモリーに100を入れる
//   var n int = 100
//   fmt.Println(n)
//
//   // メモリのアドレスが表示される
//   fmt.Println(&n)
//
//   // &（アンパサンド）はAddressを指す
//   // ポインタを入れる型は *int
//   var p *int = &n
//   fmt.Println(p)
//
//   // *（アスタリスク）はメモリの中身の実態を参照
//   // 100
//   fmt.Println(*p)
}