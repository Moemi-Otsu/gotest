package main

import (
  "fmt"
  "time"
)

func getOsName() string{
  return "asdfasdf"
}

func main(){
  switch os := getOsName(); os {
    case "mac":
      fmt.Println("Mac!!")
    case "windows":
      fmt.Println("Windows!!")
    default:
      fmt.Println("Default!!", os)
  }
//   fmt.Println(os)

  // timeを使用
  // goではオブジェクトではなくストラクトと呼ぶ
  t := time.Now()
  fmt.Println(t.Hour())
  switch {
    case t.Hour() < 12:
      fmt.Println("Morning")
    case t.Hour() < 17:
      fmt.Println("Afternoon")
    case t.Hour() < 24:
      fmt.Println("Evening")
  }
}