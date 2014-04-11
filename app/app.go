package app

import (
  "errors"
  "fmt"
  "github.com/stretchr/goweb"
  "github.com/stretchr/goweb/context"
  "github.com/tgreiser/testapp"
  "github.com/tgreiser/testapp/pinger"
  "net/http"
  "os"
)

func Init() {
  var _ = pinger.Ping()
}

func init() {
  testapp.MapRoutes()
  http.Handle("/", goweb.DefaultHttpHandler())
}


