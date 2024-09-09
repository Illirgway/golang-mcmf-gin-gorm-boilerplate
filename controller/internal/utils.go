/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package internal

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type PageField struct {
	Page uint `form:"page" json:"page" binding:"max=255"`
}

func ShouldBindForm(ctx *gin.Context, obj any) error {
	return ctx.ShouldBindWith(obj, binding.Form)
}

func ShouldBindFormPost(ctx *gin.Context, obj any) error {
	return ctx.ShouldBindWith(obj, binding.FormPost)
}

// SEE gin.Context::ShouldBind()
func ShouldBindFormAny(ctx *gin.Context, obj any) error {

	var b binding.Binding

	switch ctx.ContentType() {
	case binding.MIMEMultipartPOSTForm:
		b = binding.FormMultipart
	case binding.MIMEPOSTForm:
		b = binding.FormPost
	default:
		b = binding.Form
	}

	return ctx.ShouldBindWith(obj, b)
}

func IsAjaxRequest(ctx *gin.Context) bool {
	h := ctx.GetHeader("X-Requested-With")
	return (h != "") && strings.EqualFold(h, "XMLHttpRequest")
}
