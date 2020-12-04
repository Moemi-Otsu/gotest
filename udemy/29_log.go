package main

import "log"

func main(){
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
}