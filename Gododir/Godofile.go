package main

import (
//	"fmt"
	. "gopkg.in/godo.v1"
)

func tasks(p *Project) {
	Env = `GOPATH=.vendor::$GOPATH`

	p.Task("buildAll", D{"golint"}, func() error {
		return Bash(`gom exec gox ./src/...`)
	})

	p.Task("golint", func() error {
		return Bash(`gom exec golint ./src/...`)
	}).Watch("**/*.go")

	p.Task("testAll", D{"golint"}, func() error {
		return Bash(`gom exec go test -v ./src/...`)
	}).Watch("**/*.go")
}

func main() {
	Godo(tasks)
}