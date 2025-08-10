package render

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/mod/modfile"
	"golang.org/x/tools/imports"
)

var (
	gomodFile = "go.mod"

	tmplPath = "tmpl"

	tmplExtDst = ".go.tmpl"

	tmplExtSrc = ".go"
)

//go:embed tmpl/* tmpl/**/*
var templatesFS embed.FS

func RenderTemplate(tpl *template.Template, sub any, filename string) ([]byte, error) {
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, sub); err != nil {
		return nil, fmt.Errorf("template.Execute failed: %w", err)
	}

	formatted, err := imports.Process("", buf.Bytes(), &imports.Options{
		Fragment: true,
	})
	if err != nil {
		return nil, fmt.Errorf("imports.Process failed: %v", err)
	}
	return formatted, nil
}

func Render(dst string) (err error) {
	m, err := getModuleName()
	if err != nil {
		return err
	}

	if dst != "" {
		m += "/" + dst
	}

	sub := map[string]string{
		"Module": m,
	}

	fmt.Println("sub:", sub)

	err = fs.WalkDir(templatesFS, tmplPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		dstPath, err := filepath.Rel(tmplPath, path)
		if err != nil {
			return err
		}

		content, err := templatesFS.ReadFile(path)
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, tmplExtDst) {
			dstPath = filepath.Join(dst, dstPath)
			dstPath = strings.Replace(dstPath, tmplExtDst, tmplExtSrc, 1)

			base := filepath.Base(path)

			tmpl, err := template.New(base).Parse(string(content))
			if err != nil {
				return err
			}

			content, err = RenderTemplate(tmpl, sub, base)
			if err != nil {
				return err
			}
		}

		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			return err
		}
		return os.WriteFile(dstPath, content, 0644)
	})

	return err

}

func getModuleName() (string, error) {
	data, err := os.ReadFile(gomodFile)
	if err != nil {
		return "", err
	}

	f, err := modfile.Parse(gomodFile, data, nil)
	if err != nil {
		return "", err
	}

	return f.Module.Mod.Path, nil
}
