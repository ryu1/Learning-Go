package main

import (
//	"fmt"
	. "gopkg.in/godo.v1"
)

func tasks(p *Project) {
	Env = `GOPATH=.vendor::$GOPATH`

//	p.Task("install", func() {
//		Run("go get github.com/golang/lint/golint")
//		Run("go get github.com/mgutz/goa")
//		Run("go get github.com/robertkrimen/godocdown/godocdown")
//	})

	p.Task("buildAll", D{"lint"}, func() error {
		return Bash(`gom exec gox ./src/...`)
	})

	p.Task("lint", func() error {
		//Bash(`gom exec golint ./src/...`)
		Run("golint .")
		Run("gofmt -w -s ./src/...")
		Run("go vet ./src/...")
		return
	}).Watch("**/*.go")

	p.Task("testAll", D{"lint"}, func() error {
		return Bash(`gom exec go test -v ./src/...`)
	}).Watch("**/*.go")
}

func main() {
	Godo(tasks)
}