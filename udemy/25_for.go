package main

import "fmt"

func main(){
  for i := 0; i < 10; i++ {
    if i == 3{
      fmt.Println("continue")
      continue
      // continueはfor文の実行をskipする
    }

    // break文
    // 特定の条件でループを終了させる
    if i > 5 {
      fmt.Println("break")
      break
    }
    fmt.Println(i)
  }

  sum := 1
  fot ; sum < 10; {
    sum += sum
  }
  fmt.Println(sum)
}