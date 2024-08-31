/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package app

import (
	"errors"
	"fmt"

	"github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/utils/thtml"
	"github.com/Masterminds/sprig/v3"
)

/*
func (app *App) initPresenter() (err error) {

	app.presenter, err = presenter.NewPresenter(app.cfg.TemplatesRoot, sprig.FuncMap())

	if err != nil {
		return fmt.Errorf("presenter creation error: %w", err)
	}

	return nil
}
*/

var (
	errNoRouter = errors.New("App template loading error: must be called after router init")
)

func (app *App) loadTemplates() (err error) {

	if app.router == nil {
		return errNoRouter
	}

	app.router.FuncMap = sprig.FuncMap()

	app.presenter, err = thtml.LoadTemplates(app.cfg.Paths.Templates, app.router.FuncMap)

	if err != nil {
		return fmt.Errorf("App template loading error: %w", err)
	}

	app.router.SetHTMLTemplate(app.presenter)

	return nil
}
