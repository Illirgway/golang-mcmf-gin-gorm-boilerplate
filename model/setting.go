/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package model

const (
	settingsTableName = "settings"
)

type SettingId uint

//go:nosplit
func (id SettingId) Raw() uint {
	return uint(id)
}

//

type SettingType string

type SettingData struct {
	Name  string `gorm:"column:name;size:31;not null;uniqueIndex:uniq_name" json:"name" form:"name" binding:"required,max=31,printascii"`
	Value string `gorm:"column:value;size:255;not null" json:"value" form:"value" binding:"omitempty,max=255,utf8-string"`
}

type SettingContent struct {
	SettingData
	Desc string `gorm:"column:desc;type:TEXT;not null;size:4096" json:"desc" form:"desc" binding:"omitempty,max=4095,utf8-text"`
}

type SettingIdField struct {
	Id SettingId `gorm:"column:id;type:INT UNSIGNED;size:10;primaryKey;autoIncrement" json:"id" form:"id" binding:"required,min=1"`
}

type Setting struct {
	SettingIdField
	SettingContent
	CommonUpdatedAtField
}

// TableName implements interface gorm.Tabler
//
//go:nosplit
func (*Setting) TableName() string {
	return settingsTableName
}

// Reset for pooling
//
//go:nosplit
func (s *Setting) Reset() {
	*s = Setting{}
}

//

type SettingsList []Setting

type SettingsKV map[string]string // name -> value
