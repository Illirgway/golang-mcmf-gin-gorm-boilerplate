/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package internal

import (
	"fmt"

	"github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	SessionKeyAdminId = "aid"
	// TODO SessionKeyIsAdmin = "is_admin"
)

func SessionAdminId(ctx *gin.Context) (id model.UserId, err error) {

	sess := sessions.Default(ctx)

	v := sess.Get(SessionKeyAdminId)

	if v == nil {
		return model.NoUser, nil
	}

	id, ok := v.(model.UserId)

	if !ok {
		return model.NoUser, fmt.Errorf("wrong session admin id value's %[1]v type: %[1]T (must be %[2]T)", v, model.NoUser)
	}

	return id, nil
}
