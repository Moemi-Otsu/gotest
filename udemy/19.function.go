package main

import "fmt"

// go の func の構文はこれ？？
// 返り値に自分で設定した変数名を使って返すこともできる。

// func 関数名(変数, 型) (返り値, 返り値の型) {
// }

// 2つ変数を返したいとき
func add(x, y int) (int, int) {
  return x + y, x - y
}

// nameのreturn values
// 返り値に変数を設定して返すことができる。
// この例ではresult
// 返り値が多くなったときに、わかりやすい
func cal(price, item int) (result int) {
  result = price * item
  return result
}

func convert(price int) float64 {
  return float64(price)
}

func main() {
  // 返り値をr1, r2 と2つ書く必要がある
  r1, r2 := add(10, 20)
  fmt.Println(r1, r2)

  r3 := cal(100, 2)
  fmt.Println(r3)

  f := func(x int) {
    fmt.Println("inner func", x)
  }
  f(1)

  // (1) は変数に入れなくても実行できる
  func(x int) {
    fmt.Println("inner func", x)
  }(1)
}