// 構造体
package main

import "fmt"

type Vertex struct{
  // 小文字で宣言するとprivateになる
  X, Y int
  S string
}

func changeVertex(v Vertex){
  v.X = 1000
}

func changeVertex2(v *Vertex){
  (*v).X = 1000
}

func main() {
  v := Vertex{1, 2, "test"}
  changeVertex(v)
  fmt.Println(v)

  v2 := &Vertex{1, 2, "test"}
  changeVertex2(v2)
  fmt.Println(v2)
  /*
  v := Vertex{X: 1, Y:2}
  fmt.Println(v)
  // {1 2 }
  fmt.Println(v.X, v.Y)
  // 1 2

  v2 := Vertex{X: 1}
  fmt.Println(v2)
  // {1 0 }

  v3 := Vertex{1, 2, "test"}
  fmt.Println(v3)
  // {1 2 test}

  v4 := Vertex{}
  fmt.Println(v4)
  // {0 0 }
  fmt.Printf("%T %v\n", v4, v4)

  var v5 Vertex
  fmt.Println(v5)
  // {0 0 }
  fmt.Printf("%T %v\n", v5, v5)

  // sliceとかmapはnilだったが、vertexはnilではない

  v6 := new(Vertex)
  fmt.Println(v6)
  // &{0 0 }
  fmt.Printf("%T %v\n", v6, v6)
  // *main.Vertex &{0 0 }

  v7 := &Vertex{}
  fmt.Printf("%T %v\n", v7, v7)
  // *main.Vertex &{0 0 }

  s := make([]int, 0)
//   s := []int{}
  fmt.Println(s)

  */
}