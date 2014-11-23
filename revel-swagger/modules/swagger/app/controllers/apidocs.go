package controllers

import (
  "strings"

  "github.com/revel/revel"
)

type Docs struct {
  *revel.Controller
}

func (c *Docs) GetApiDocs() revel.Result {
  isJsonRequest := false

  if acceptHeaders, ok := c.Request.Header["Accept"]; ok {
    for _, acceptHeader := range acceptHeaders {
      if strings.Contains(acceptHeader, "json") {
        isJsonRequest = true
        break
      }
    }
  }

  if isJsonRequest {
    return c.RenderText(resourceListingJson)
  } else {
    return c.Redirect("/docs/api/index.html")
  }
}

func (c *Docs) GetApiDoc() revel.Result {
  apiKey := strings.TrimPrefix(c.Request.RequestURI, "/docs/")
  apiKey = strings.TrimSuffix(apiKey, "/")

  if json, ok := apiDescriptionsJson[apiKey]; ok {
    return c.RenderText(json)
  } else {
    return c.NotFound("404")
  }
}
