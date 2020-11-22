package main

import "fmt"

// 返り値に ()func() int) を入れる
func incrementGenerator() (func() int){
  x := 0
  return func() int{
    x++
    return x
  }
}

// どういった時にクロージャーを使うのか
func circleArea(pi float64) func(radius float64)  float64{
  return func(radius float64) float64{
    return pi * radius * radius
  }
}

func main() {
  // このcounterには、
  // x := 0...部分が入っている
  counter := incrementGenerator()
  fmt.Println(counter())
  fmt.Println(counter())
  fmt.Println(counter())
  fmt.Println(counter())

  c1 := circleArea(3.14)
  fmt.Println(c1(2))

  c2 := circleArea(3)
  fmt.Println(c2(2))
}