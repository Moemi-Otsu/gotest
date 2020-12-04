package main

import (
  "log"
  "fmt"
  "os"
  "io"
)

func LoggingSettings(logFile string){
  // パイプ|を使って複数の条件を書ける
  // 0666はパーミッション
  logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  multiLogFile := io.MultiWriter(os.Stdout, logfile)
  log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
  log.SetOutput(multiLogFile)
}

func main(){
  LoggingSettings("test.log")
  _, err := os.Open("fdsfdfds")
  if err != nil{
    log.Fatalln("Exit", err)
  }

  /*
  他の言語では、loggingライブラリにこんなのあったりする
  でも、goにはこれらは存在しない
  サードパーティのライブラリを使う必要がある。
  logging.debug("")
  logging.info("")
  logging.warn("")
  logging.error("")
  logging.exception("")
  */

  log.Println("logging!")
  log.Printf("%T %v", "test", "test")

  // フェイタルを使うと、その時点でコードが終了する
  log.Fatalf("%T %v", "test", "test")
  log.Fatalln("error!!")

  // 実行されない
  fmt.Println("ok!")
}