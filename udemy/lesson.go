package main

import "fmt"

var (
  i int = 1
  f64 float64 = 1.2
  s string = "test"
  t, f bool = true, false
)

func foo() {
  xi := 1
  xf64 := "test"
  xt, xf := true, false
  fmt.Println(xi, xf64, xt, xf)
  fmt.Printf("%T", xf64)
 }

func main() {
  fmt.Println(i, f64, s, t, f)
  foo()
}