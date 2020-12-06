// goでは、panicを使うことは推奨していない。
// 前の章でやったエラーをしっかり認識してエラーとして返してやることを推奨している。
package main

import "fmt"

// DBにコネクションするようなパッケージを使っているとする
func thirdPartyConnectDB(){
  panic("Unable to connect database!")
}

func save(){
  // deferで、saveが実行された後に
  // panicを起こしたものをrecoverでキャッチできる
  // キャッチした物をsに格納してPrintしている
  // recoverで、システムを強制終了させないようpanicを回避できる
  defer func(){
    s := recover()
    fmt.Println(s)
  }()
  thirdPartyConnectDB()
}

func main() {
  save()
  fmt.Println("OK?")
}