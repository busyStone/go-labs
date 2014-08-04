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
  SliceLength = 100000
)

func getStructsSlice() []T {
  t := []T{}

  for i := 0; i < SliceLength; i++ {
    t = append(t, T{Name: "test" + fmt.Sprintf("%d", i)})
  }

  return t
}

func getFixedStructsSlice() []T {
  t := make([]T, SliceLength)

  for i := 0; i < SliceLength; i++ {
    t = append(t, T{Name: "test" + fmt.Sprintf("%d", i)})
  }

  return t
}

func getStructPointersSlice() []*T {
  t := []*T{}

  for i := 0; i < SliceLength; i++ {
    t = append(t, &T{Name: "test" + fmt.Sprintf("%d", i)})
  }

  return t
}

func getFixedStructPointersSlice() []*T {
  t := make([]*T, SliceLength)

  for i := 0; i < SliceLength; i++ {
    t = append(t, &T{Name: "test" + fmt.Sprintf("%d", i)})
  }

  return t
}

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())

  getStructsSlice()

  ch := make(chan os.Signal)
  signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
  log.Println(<-ch)
}
