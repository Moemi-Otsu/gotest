package main

import "fmt"

func by2(num int) string{
  if num % 2 == 0 {
    return "ok"
  }else{
    return "no"
  }
}

func main() {

  // この書き方は、
  result := by2(10)
  if result == "ok" {
    fmt.Println("great")
  }
  fmt.Println(result)

  // この書き方にもできる。
  if result2 := by2(10); result2 == "ok"{
    fmt.Println("great 2")
  }

//   num := 6
//   if num % 2 == 0 {
//     fmt.Println("by 2")
//   }else if num % 3 == 0 {
//     fmt.Println("by 3")
//   }else{
//     fmt.Println("else")
//   }

  x, y := 10, 10
  if x == 10 && y == 10 {
    fmt.Println("&&")
  }

  if x == 10 || y == 10 {
    fmt.Println("||")
  }
}