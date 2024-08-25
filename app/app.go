/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package app

import (
	"github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/config"
	"github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/proto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	cfg      *config.Config
	db       *gorm.DB
	router   *gin.Engine
	repos    repos
	services services
}

//go:nosplit
func (app *App) DB() *gorm.DB {
	return app.db
}

//go:nosplit
func (app *App) Config() *config.Config {
	return app.cfg
}

//go:nosplit
func (app *App) Router() *gin.Engine {
	return app.router
}

//go:nosplit
func (app *App) Repositories() proto.Repositories {
	return &app.repos
}

//go:nosplit
func (app *App) Services() proto.Services {
	return &app.services
}

func New(cfg *config.Config) *App {
	return &App{
		cfg: cfg,
	}
}
