package main

import (
  "log"
  "os"
  "os/signal"
  "runtime"
  "syscall"
)

type T struct {
  Hold []string
}

const (
  SliceLen = 10000
)

func newT() (*T, error) {
  hold := make([]string, SliceLen)
  return &T{hold}, nil
}

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())

  var err error
  for i := 0; i < SliceLen; i++ {
    _, err = newT()
    if err != nil {
      break
    }
  }

  ch := make(chan os.Signal)
  signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
  log.Println(<-ch)
}
