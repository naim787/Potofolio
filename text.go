package main 
import (
  "fmt"
  )

func AddData(name string, channel chan string) {
  channel <-name
}

func main() {
  channel := make(chan string)
  AddData("naim", cahnnel)
  fmt.Println(channel)
  fmt.Println("tino")
  }
