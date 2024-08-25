/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package model

// TODO

type UserId uint

//go:nosplit
func (id UserId) Raw() uint {
	return uint(id)
}

const NoUser = UserId(0)
