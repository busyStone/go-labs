package main

import (
  "fmt"
  "log"
  "os"
  "os/signal"
  "runtime"
  "syscall"
)

type T struct {
  Name string
}

const (
  SliceLen = 10000
)

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())

  s := make([]T, SliceLen)

  for i := 0; i < SliceLen; i++ {
    s = append(s, T{Name: fmt.Sprintf("%d", i)})
  }

  ch := make(chan os.Signal)
  signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
  log.Println(<-ch)
}
