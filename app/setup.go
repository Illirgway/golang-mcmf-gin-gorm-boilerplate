/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package app

import (
	"errors"
	"net/http"

	"github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var (
	errAlreadySetup = errors.New("app is already setup, called twice")
)

func (app *App) Setup() (err error) {

	if app.router != nil {
		return errAlreadySetup
	}

	/*
		if err = app.initPresenter(); err != nil {
			return err
		}
	*/

	if err = app.initAppDb(); err != nil {
		return err
	}

	if err = app.initRepos(); err != nil {
		return err
	}

	if err = app.initServices(); err != nil {
		return err
	}

	// ATN! между сервисами и роутингом
	if err = utils.InitValidators(); err != nil {
		return err
	}

	if err = app.initRouter(); err != nil {
		return err
	}

	return app.loadTemplates()
}

func (app *App) initRouter() (err error) {

	r := gin.Default()

	// SEE https://stackoverflow.com/questions/74592358/getting-the-remote-ip-address-when-using-nginx-proxy-for-glang-gin
	r.TrustedPlatform = "X-Real-IP"

	// SEE https://pkg.go.dev/github.com/gin-gonic/gin#Engine.SetTrustedProxies
	if err = r.SetTrustedProxies(nil); err != nil {
		return err
	}

	// TODO ginpprof if debug

	// session store
	{
		store := cookie.NewStore([]byte(app.cfg.Sessions.Secret))

		timeout := app.cfg.Sessions.Timeout

		if timeout <= 0 {
			timeout = 86400 * 30 // default gorilla timeout
		}

		store.Options(sessions.Options{
			Path:     "/",
			MaxAge:   int(timeout),
			Secure:   true, // TODO? сделать зависимым от https бекенда?
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})

		r.Use(sessions.Sessions("sess", store))
	}

	app.router = r

	return nil
}
