package main

import (
  "fmt"
  "os"
)

func foo(){
  defer fmt.Println("world foo")

  fmt.Println("hello foo")
}

func main(){
  foo()

  // deferは遅延実行の意味
  // main関数を実行した後にdefer処理が実行される。
  defer fmt.Println("world")
  fmt.Println("hello")

  // スタッキングデファー
  // はじめに入れたものが最後に実行される
  // run, success, 3, 2, 1 の順番
  fmt.Println("run")
  defer fmt.Println(1)
  defer fmt.Println(2)
  defer fmt.Println(3)
  fmt.Println("success")

  // deferをどういうときに使うか
  // fileをopenするやり方で実践
  file, _ := os.Open("./lesson.go")
  // fileは必ずcloseしてやる必要がある
  // deferにしておけば、実行順序を守りつつ忘れる心配なし
  defer file.Close()
  data := make([]byte, 100)
  file.Read(data)
  fmt.Println(string(data))
}