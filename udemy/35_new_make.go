package main

import "fmt"

func main() {
  // 空のslice
  // make を使う
  s := make([]int, 0)
  fmt.Printf("%T\n", s)

  // 空のmap
  // make を使う
  m := make(map[string]int)
  fmt.Printf("%T\n", m)

  // 空のポインタ
  // ポインタの場合は new を使う
  var p *int = new(int)
  fmt.Printf("%T\n", p)

  // ポインタを返すか返さないかで、newとmakeを使い分ける

  /*
  // new
  // 値を何も入れない状態でメモリにポインタが入る領域を確保したい場合
  // 0xc000014078
  var p *int = new(int)
  fmt.Println(p)
  *p++ //1を足す
  fmt.Println(*p)

  // 値がまだできていない状態
  // <nil>
  var p2 *int
  fmt.Println(p2)
  *p2++
  fmt.Println(p2)
  */
}