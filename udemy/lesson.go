package main

import "fmt"

// Initは特別な関数名。
// main関数よりも先に呼ばれる。
func init() {
  fmt.Println("Init!")
}

// オリジナル関数。このままだと実行されない。
// main 関数の中で呼んであげる必要がある。
func bazz() {
  fmt.Println("Bazz")
}

func main() {
  bazz()
  // fmt はフォーマットというライブラリ
  fmt.Println("Hello World!", "TEST TEST")
}