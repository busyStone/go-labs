package main

import (
  "log"
  "net/http"

  "github.com/Unknwon/macaron"
)

func main() {
  m := macaron.Classic()

  m.Get("/", mHandler)

  log.Println("sever is running...")

  log.Println(http.ListenAndServe("0.0.0.0:4000", m))
}

func mHandler(ctx *macaron.Context) string {
  return "the request path is: " + ctx.Req.RequestURI
}
