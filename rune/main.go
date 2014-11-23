package main

import (
  "fmt"
  "log"
  "strings"
)

func main() {
  str1 := "Hello 世界！"

  fmt.Println(str1)

  log.Println(len(str1))

  log.Println(len([]rune(str1)))

  str2 := " " + str1 + " "

  trim(str2)

  log.Println(str2)

}

func trim(text string) {

  text = strings.TrimSpace(text)

  log.Println(text)
}
