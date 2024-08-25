/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package model

import "time"

type CommonUpdatedAtField struct {
	UpdatedAt time.Time `gorm:"column:updated_at;type:INT UNSIGNED;size:10;not null;autoUpdateTime" json:"mtime"`
}
