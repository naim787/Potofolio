package main 
import (
  "fmt"
  )

func main() {
  channel := make(chan string)
  channel <- "naim mulai menjadi masuk arena";
  fmt.Println(<-channnel)
  }
