package views

import (
	"fmt"
	"io"

	"github.com/realistschuckle/gohaml"
)

func WriteStatic(path string, w io.Writer) error {
	data, err := Asset(path)
	if err != nil {
		return err
	}
	w.Write(data)
	return nil
}

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

func (self *Engine) Render(w io.Writer) error {
	data, err := Asset(self.path)
	if err != nil {
		return err
	}

	engine, err := gohaml.NewEngine(string(data))
	if err != nil {
		return err
	}

	output := engine.Render(self.scope)

	fmt.Fprint(w, output)
	return nil
}
