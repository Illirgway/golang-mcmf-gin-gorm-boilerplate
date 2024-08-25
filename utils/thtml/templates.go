/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package thtml

import (
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// SEE https://stackoverflow.com/questions/38686583/golang-parse-all-templates-in-directory-and-subdirectories

type TemplateValues map[string]any

const (
	Sep = '/'
	Ext = "thtml"
)

type walker struct {
	dir   string
	funcs template.FuncMap
	t     *template.Template
	sep   bool
}

// SEE https://stackoverflow.com/questions/20716726/call-other-templates-with-dynamic-name
// TODO уродливо. рефакторинг
type callTemplateExecutor struct {
	t *template.Template
}

func (cte *callTemplateExecutor) CallTemplate(name string, data interface{}) (ret template.HTML, err error) {

	// TODO вести буфер оптимальныйх развмеров в виде map: name -> size (а можно еще зависимость от data вычислять)

	// WARN лучше, чем bytes.Buffer (минимум на 1 мемаллок/мемкопи меньше), но нельзя переиспользовать (пулить) сам Builder!
	buf := new(strings.Builder)

	const initialBuffer = 2 << 10 // на глаз стартуем с 2 Кб

	// TODO предварительный расчет для Grow
	buf.Grow(initialBuffer)

	if err = cte.t.ExecuteTemplate(buf, name, data); err != nil {
		return "", err
	}

	return template.HTML(buf.String()), nil
}

func attachCallTemplateExecutor(t *template.Template) {

	e := &callTemplateExecutor{t}

	t.Funcs(template.FuncMap{
		"CallTemplate": e.CallTemplate,
	})
}

var (
	sepStr   = string(Sep)
	fsSepStr = string(filepath.Separator)
)

func (w *walker) name(path string) (name string, err error) {

	rel, err := filepath.Rel(w.dir, path)

	if err != nil {
		return "", err
	}

	if w.sep {
		// fsPathSeparator -> TemplateNameSepStr
		rel = strings.ReplaceAll(rel, fsSepStr, sepStr)
	}

	// name = rel - ext
	name = rel[:len(rel)-(1+len(Ext))]

	return name, nil
}

func (w *walker) addFile(path string, name string) (err error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	_, err = w.t.New(name).Parse(string(data))

	return err
}

func (w *walker) walkFn(path string, info fs.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if !info.Mode().IsRegular() {
		return nil
	}

	if ext := filepath.Ext(path); len(ext) <= 0 || (ext[1:] != Ext) {
		return nil
	}

	name, err := w.name(path)

	if err != nil {
		return err
	}

	//
	return w.addFile(path, name)
}

func (w *walker) load() error {

	w.sep = filepath.Separator != Sep

	w.t = template.New("")

	if len(w.funcs) > 0 {
		w.t.Funcs(w.funcs)
	}

	// SEE https://stackoverflow.com/questions/20716726/call-other-templates-with-dynamic-name
	// TODO уродливо. рефакторинг
	attachCallTemplateExecutor(w.t)

	return filepath.Walk(w.dir, w.walkFn)
}

func LoadTemplates(rootDir string, funcs template.FuncMap) (t *template.Template, err error) {

	if rootDir, err = filepath.Abs(rootDir); err != nil {
		return nil, err
	}

	w := walker{
		dir:   rootDir,
		funcs: funcs,
	}

	if err = w.load(); err != nil {
		return nil, err
	}

	return w.t, nil
}
