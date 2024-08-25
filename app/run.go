/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package app

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
)

var (
	errAppUnsuspectedRun = errors.New("App.Run should be called AFTER App.Setup, but called before")
)

func (app *App) Run() (err error) {

	if app.router == nil {
		return errAppUnsuspectedRun
	}

	if err = app.repos.run(); err != nil {
		return fmt.Errorf("repositories starting error: %w", err)
	}

	if err = app.services.run(); err != nil {
		return fmt.Errorf("services starting error: %w", err)
	}

	// force free unneeded mem before run
	runtime.GC()

	err = app.router.Run(app.cfg.ListenAddr())

	if err == http.ErrServerClosed {
		err = nil
	}

	err1 := app.stop()

	if err == nil {
		err = err1
	}

	return err
}

func (app *App) stop() (err error) {

	if err = app.repos.stop(); err != nil {
		return err
	}

	if err = app.services.stop(); err != nil {
		return err
	}

	db, err := app.db.DB()

	if err != nil {
		return err
	}

	return db.Close()
}
