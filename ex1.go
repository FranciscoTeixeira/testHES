package main

  import "fmt"

  func main() {

    for i:=12; i >= 0; i-- {

      if (i == 2 || i == 4 || i == 10) {
        continue
      }

      fmt.Println("i =", i)
      
    }

  }
