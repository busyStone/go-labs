package main

import (
  "container/list"
  "log"
)

func main() {

  l := list.New()

  l.PushBack(8)

  e := l.Front()

  log.Println(e.Value)

  l.Remove(e)

  log.Println(l.Front())
}
