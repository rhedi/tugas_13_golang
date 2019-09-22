package main

import "fmt"
import "encoding/base64"
import "crypto/sha1"

func main()  {
  var nama = "rhedi pratama"

    var rahasia = base64.StdEncoding.EncodeToString([]byte(nama))
    fmt.Println("Nama Saya Adalah", rahasia)

    var dibuka,_ = base64.StdEncoding.DecodeString("cmhlZGkgcHJhdGFtYQ==")
    var codedibuka = string(dibuka)
    fmt.Println("nama saya adalah", codedibuka)

    var secret = sha1.New()
    secret.Write([]byte(nama))
    var enkrip = secret.Sum(nil)
    var dikunci = fmt.Sprintf("%x", enkrip)

    fmt.Println(dikunci)
}
