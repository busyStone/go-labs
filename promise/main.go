package main

import (
  "log"
  "net/http"
  "time"

  "github.com/fanliao/go-promise"
)

func main() {
  task := func() (interface{}, error, bool) {
    url := "http://example.com/"

    resp, err := http.Get(url)
    defer resp.Body.Close()
    if err != nil {

      return url, err, false
    }
    return url, resp.Body, true
  }

  f := promise.Start(task).OnSuccess(func(r interface{}) {
    log.Println("body ", r)
  }).OnFailure(func(e interface{}) {
    log.Println("error ", e.(error))
  }).OnComplete(func(v interface{}) {
    log.Println("complete ", v)
  })

  time.Sleep(10 * time.Second)
}
