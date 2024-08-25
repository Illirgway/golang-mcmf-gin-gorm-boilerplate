/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package controller

import (
	"net/http"

	"github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/controller/internal"
	"github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/model"
	"github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/proto"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	settings proto.RepositorySettings
}

func (c *Controller) Before(ctx *gin.Context) {

	id, err := internal.SessionAdminId(ctx)

	if err != nil {
		_ = internal.AbortWithServerError(ctx, err)
		return
	}

	// авторедиректим в панель если залогинены
	if path := ctx.Request.URL.Path; (id != model.NoUser) && (path != "") && (path != "/") && (path != "/logout") {
		internal.RedirectWithAbort(ctx, http.StatusTemporaryRedirect, "/")
	}
}

func (c *Controller) Get(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "frontend/index", nil)
}

func NewFrontend(repos proto.Repositories) *Controller {
	return &Controller{
		settings: repos.Settings(),
	}
}
