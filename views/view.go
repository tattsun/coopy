package views

import (
	"fmt"
	"io"

	"github.com/realistschuckle/gohaml"
)

type Engine struct {
	path  string
	scope map[string]interface{}
}

func NewEngine(path string) *Engine {
	scope := make(map[string]interface{})
	return &Engine{path: path, scope: scope}
}

func (self *Engine) Add(key string, val interface{}) {
	self.scope[key] = val
}

func (self *Engine) Render(w io.Writer) {
	data, err := Asset(self.path)
	if err != nil {
		panic(err)
	}

	engine, err := gohaml.NewEngine(string(data))
	if err != nil {
		panic(err)
	}

	output := engine.Render(self.scope)

	fmt.Fprint(w, output)
}
