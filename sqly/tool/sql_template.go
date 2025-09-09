package tool

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/llyb120/bingo/config"
	"github.com/llyb120/bingo/core"
	"github.com/llyb120/bingo/log"
	"github.com/llyb120/gotemplate"
)

var GoTemplateStarter core.Starter = func() func() {
	var cfg = core.Require[config.Config]()
	var render = gotemplate.NewSqlRender()
	path := cfg.GetString("sql.template.path")
	filesName := cfg.GetString("sql.template.files")
	env := cfg.GetString("server.environment")
	if env == "dev" {
		if err := startDev(render, path); err != nil {
			panic(err)
		}
	} else {
		if filesName == "" {
			panic("sql.template.files is empty")
		}
		embeddedFiles := core.Require[embed.FS](filesName)
		if err := startProd(render, embeddedFiles); err != nil {
			panic(err)
		}
	}
	core.ExportInstance(render, core.RegisterOption{Name: "sql-template"})
	return nil
}

func GetSql(path string, data ...any) (sql string, params []any, err error) {
	var main string
	var sub string
	pos := strings.Index(path, ".")
	if pos == -1 {
		return "", nil, fmt.Errorf("sql path is not valid: %s", path)
	}
	main = path[:pos]
	sub = path[pos+1:]
	var d any
	if len(data) > 0 {
		d = data[0]
	}
	return core.Use[gotemplate.SqlRender]().GetSql(main, sub, d)
}

func startDev(render *gotemplate.SqlRender, dir string) error {
	var g gotemplate.ErrGroup
	if err := render.Scan(func(handler gotemplate.ScanHandler) error {
		var files []fs.DirEntry
		var err error
		files, err = os.ReadDir(dir)
		if err != nil {
			return err
		}
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if !strings.HasSuffix(file.Name(), ".md") {
				continue
			}
			log.Info(nil, "scan file: %s", file.Name())
			fileName := file.Name()
			g.Go(func() error {
				var content []byte
				var err error
				content, err = os.ReadFile(filepath.Join(dir, fileName))
				if err != nil {
					return err
				}
				return handler(fileName, string(content))
			})
		}
		return nil
	}); err != nil {
		return err
	}
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func startProd(render *gotemplate.SqlRender, embeddedFiles *embed.FS) error {
	var g gotemplate.ErrGroup
	if err := render.Scan(func(handler gotemplate.ScanHandler) error {
		var files []fs.DirEntry
		var err error
		files, err = embeddedFiles.ReadDir(".")
		if err != nil {
			return err
		}
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if !strings.HasSuffix(file.Name(), ".md") {
				continue
			}
			log.Info(nil, "scan file: %s", file.Name())
			fileName := file.Name()
			g.Go(func() error {
				var content []byte
				var err error
				content, err = embeddedFiles.ReadFile(fileName)
				if err != nil {
					return err
				}
				return handler(fileName, string(content))
			})
		}
		return nil
	}); err != nil {
		return err
	}
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}
