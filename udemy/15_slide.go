package main

import "fmt"

func main() {
  n := []int{1, 2, 3, 4, 5, 6}

  // 配列のすべて
  fmt.Println(n)

  // 配列の〜2番目
  fmt.Println(n[2])

  // 配列の2〜4番目まで
  fmt.Println(n[2:4])

  // 配列の〜2番目まで
  fmt.Println(n[:2])

  // 配列の2番目以降
  fmt.Println(n[2:])

  // 初めから最後まで
  fmt.Println(n[:])

  // [1 2 100 4 5 6]
  n[2] = 100
  fmt.Println(n)

  // [[0 1 2] [3 4 5] [6 7 8]]
  var board = [][]int{
    []int{0, 1, 2},
    []int{3, 4, 5},
    []int{6, 7, 8},
  }
  fmt.Println(board)

  // [1 2 100 4 5 6 100]
  n = append(n, 100, 200, 300, 400)
  fmt.Println(n)
}