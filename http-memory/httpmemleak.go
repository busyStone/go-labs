package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "os/signal"
  "runtime"
  "runtime/debug"
  "syscall"
  "time"
)

// func init() {
//   http.HandleFunc("/", handler)
// }

func handler(w http.ResponseWriter, r *http.Request) {
  outStr := make([]string, 500*1024)
  outStr[1] = "l"
  fmt.Fprint(w, "Hello, world!")
}

func StartGC() {
  for {
    time.Sleep(10 * time.Second)
    runtime.GC()
    debug.FreeOSMemory()
    log.Println("Current Routines", runtime.NumGoroutine())
  }
}

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())
  go StartGC()
  http.HandleFunc("/hello", handler)
  go http.ListenAndServe(":1234", nil)

  ch := make(chan os.Signal)
  signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
  log.Println(<-ch)
}
