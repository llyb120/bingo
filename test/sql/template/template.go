package template

import (
	"embed"

	"github.com/llyb120/bingo/core"
)

//go:embed *.md
var embeddedFiles embed.FS

var TestTemplateStarter core.Starter = func() func() {
	core.ExportInstance(embeddedFiles, core.RegisterOption{Name: "test-template"})
	return nil
}
