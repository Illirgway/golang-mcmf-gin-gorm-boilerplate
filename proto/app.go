/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package proto

import "github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/model"

//

type Repository interface {
	Migrate() error
}

type RepositorySettings interface {
	Repository
	ById(id model.SettingId) (value string, exists bool)
	ByName(name string) (value string, exists bool)
	Enum() (list model.SettingsList, err error)
	EnumDirect() (model.SettingsList, error)
	EnumSynopsis() (list []model.SettingData, err error)
	EnumKV() (kv model.SettingsKV, err error)
	SaveKV(kv model.SettingsKV) error
}

//

type Repositories interface {
	Settings() RepositorySettings
}

type Services interface {
	// Service() ServiceService
}
