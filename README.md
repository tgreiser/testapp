## Code organization in appengine

This is a work in progress..

Organizing your Go/appengine project:

1. Use GOPATH
2. Use the full path for your modules (ie import "github.com/tgreiser/testapp")
3. Make a separate directory for your app.yaml and a go file that sets up your routes and some dummy imports (see: app/app.go)
