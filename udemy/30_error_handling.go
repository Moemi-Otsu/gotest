package main

import {
  "log"
  "os"
}

func main() {
  file, err := os.Open("./lesson.go")
  if err != nil{
    log.Fatal("Error!")
  }
  defer file.Close()
  data := make([]byte, 100)
  count, err := file.Read(date)
  if err != nil{
    log.Fatalln("Error")
  }
}