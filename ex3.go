package main

  import (
    "fmt"
    "math/rand"
  //  "strconv"
    "time"
  )

  func init() {
    rand.Seed(time.Now().UnixNano())
  }

  func main(){
    fmt.Println(rand.Intn(999))
  }
