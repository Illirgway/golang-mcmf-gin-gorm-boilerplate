/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package internal

import (
	"net/http"

	"clevergo.tech/jsend"
	"github.com/gin-gonic/gin"
)

var (
	JSendSuccess  = jsend.New(nil)
	JSendFail     = jsend.NewFail(nil)
	JSendSrvError = jsend.NewError("Internal server error", 0, nil)
)

func FailBadRequestAjax(ctx *gin.Context, err error) error {
	ctx.JSON(http.StatusBadRequest, &JSendFail)
	return ctx.Error(err)
}

func AbortWithServerError(ctx *gin.Context, err error) error {
	return ctx.AbortWithError(http.StatusInternalServerError, err)
}

func AbortWithServerErrorAjax(ctx *gin.Context, err error) error {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, &JSendSrvError)
	return ctx.Error(err)
}

func RedirectWithAbort(ctx *gin.Context, code int, location string) {
	ctx.Redirect(code, location)
	ctx.Abort()
}
