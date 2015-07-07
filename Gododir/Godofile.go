package main

import (
//	"fmt"
	. "gopkg.in/godo.v1"
)

func tasks(p *Project) {
	Env = `GOPATH=.vendor::$GOPATH`

	p.Task("buildAll", func() error {
		return Bash(`gom exec gox ./src/... `)
	})
}

func main() {
	Godo(tasks)
}