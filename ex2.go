package main
  import (
    "fmt"
    "strings"
    )

  func main() {
    words := [...]string{"Gate", "Comet", "Pizza"}

    for i := range words{
        if strings.Contains(words[i], "e"){
          fmt.Println("Found one!", string([]rune(words[i])[0]))
        }
    }

  }
