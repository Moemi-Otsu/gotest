package main

import "fmt"

// []byte ブラケットバイト
func main() {
  // アスキーコードで72:H 73:I
  // stringでキャストするとHIになる
  b := []byte{72, 73}
  fmt.Println(b)
  fmt.Println(string(b))

  // byteでキャストするとだと72 73 になる
  c := []byte("HI")
  fmt.Println(c)
}