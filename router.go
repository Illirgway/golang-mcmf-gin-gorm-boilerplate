/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Illirgway/go-ginext"
	"github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/app"
	"github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/controller"
)

const (
	assetsPathPrefix = "/assets/"
)

func registerControllers(app *app.App) (err error) {

	router := app.Router()

	// Assets if any
	// TODO? перенести это в App?
	if assetsDir := app.Config().Dirs.Assets; assetsDir != "" {

		if assetsDir, err = filepath.Abs(assetsDir); err != nil {
			return fmt.Errorf("assets dir path error: %w", err)
		}

		var info os.FileInfo

		info, err = os.Stat(assetsDir)

		if err != nil {
			return fmt.Errorf("assets dir stat error: %w", err)
		}

		if !info.IsDir() {
			return fmt.Errorf("assets dir %q is not directory: %v", assetsDir, info.Mode().Type())
		}

		router.Static(assetsPathPrefix, assetsDir)
	}

	repos := app.Repositories()

	ginext.AppendTrailingSlash(false)

	// front
	{
		c := controller.NewFrontend(repos)

		if err = ginext.EmbedController(router, c); err != nil {
			return err
		}
	}

	// ...

	return nil
}
