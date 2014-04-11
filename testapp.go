package testapp

import (
  "errors"
  "fmt"
  "github.com/stretchr/goweb"
  "github.com/stretchr/goweb/context"
  "github.com/tgreiser/testapp/pinger"
  "net/http"
  "os"
)

func mapRoutes() {
  goweb.Map("/test", func(c context.Context) error {
    return goweb.Respond.With(c, 200, []byte("Welcome to the Goweb example app - Ping? " + Ping()))
  })

  goweb.Map("GET", "people/me", func(c context.Context) error {
    hostname, _ := os.Hostname()
    return goweb.Respond.WithRedirect(c, fmt.Sprintf("/people/%s", hostname))
  })

  goweb.Map(func (c context.Context) error {
    var errs []error
    errs = append(errs, errors.New("Path not found"))
    return goweb.API.Respond(c, 404, nil, []string {"Path not found"} )
  })
}

func init() {
  mapRoutes()
  http.Handle("/", goweb.DefaultHttpHandler())
}
