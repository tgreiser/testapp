package app

import (
  "github.com/stretchr/goweb"
  "github.com/tgreiser/testapp"
  "github.com/tgreiser/testapp/pinger"
  "net/http"
)

func Init() {
  var _ = pinger.Ping()
}

func init() {
  testapp.MapRoutes()
  http.Handle("/", goweb.DefaultHttpHandler())
}
