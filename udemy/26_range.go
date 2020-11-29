package main

import "fmt"

func main(){
  l := []string{"python", "go", "java"}

  // for文
  for i := 0; i < len(l); i++{
    fmt.Println(i, l[i])
  }

  // 上と同じ実行結果
  for i, v := range l{
    fmt.Println(i, v)
  }

  // index番号を使わない場合は _ でスキップできる
  for _, v := range l {
    fmt.Println(v)
  }

  // mapでもできる
  m := map[string]int{"apple": 100, "banana": 200}
  for k, v := range m{
    fmt.Println(k, v)
  }
  // mapでkeyだけ取り出す場合
  for k := range m{
    fmt.Println(k)
  }
  // mapでvalueだけ取り出す場合
  for _, v := range m{
    fmt.Println(v)
  }
}