/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package repository

import (
	"errors"
	"sync"

	"github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/model"
	"gorm.io/gorm"
)

// fully cached data

type (
	settingsNamesIndex = map[string]uint // name -> index in SettingsList (NOT SETTING ID!!!)
)

type settingsCache struct {
	list  model.SettingsList
	index entityIndex
	names settingsNamesIndex
}

type Settings struct {
	db    *gorm.DB
	lock  sync.RWMutex
	cache settingsCache
}

var (
	errSettingFetchEmpty = errors.New("empty settings table")
)

func (r *Settings) ById(id model.SettingId) (value string, exists bool) {

	r.lock.RLock()
	defer r.lock.RUnlock()

	idx, exists := r.cache.index[id.Raw()]

	if !exists {
		return "", false
	}

	return r.cache.list[idx].Value, true
}

func (r *Settings) ByName(name string) (value string, exists bool) {

	r.lock.RLock()
	defer r.lock.RUnlock()

	idx, exists := r.cache.names[name]

	if !exists {
		return "", false
	}

	return r.cache.list[idx].Value, true
}

func (r *Settings) EnumSynopsis() (list []model.SettingData, err error) {

	r.lock.RLock()
	defer r.lock.RUnlock()

	src := r.cache.list

	list = make([]model.SettingData, len(src))

	for i := 0; i < len(src); i++ {
		list[i] = src[i].SettingData
	}

	return list, nil
}

func (r *Settings) EnumKV() (kv model.SettingsKV, err error) {

	r.lock.RLock()
	defer r.lock.RUnlock()

	list := r.cache.list

	kv = make(model.SettingsKV, len(list))

	for i := 0; i < len(list); i++ {
		kv[list[i].Name] = list[i].Value
	}

	return kv, nil
}

// copy cache list
func (r *Settings) Enum() (list model.SettingsList, err error) {

	r.lock.RLock()
	defer r.lock.RUnlock()

	// compiler optimization CL#146719 make+copy pattern
	// https://go-review.googlesource.com/c/go/+/146719/
	list = make(model.SettingsList, len(r.cache.list))
	copy(list, r.cache.list)

	return list, nil
}

func (r *Settings) EnumDirect() (model.SettingsList, error) {

	r.lock.RLock()
	defer r.lock.RUnlock()

	return r.cache.list, nil
}

func (r *Settings) SaveKV(kv model.SettingsKV) error {

	r.lock.Lock()
	defer r.lock.Unlock()

	vals := make(model.SettingsList, 0, len(kv))

	list, names := r.cache.list, r.cache.names

	// params : value to id : value
	for param, value := range kv {

		idx := names[param]

		v := model.Setting{
			SettingIdField: model.SettingIdField{
				Id: list[idx].Id,
			},
			SettingContent: model.SettingContent{
				SettingData: model.SettingData{
					Value: value,
				},
			},
		}

		vals = append(vals, v)
	}

	tx := r.db.Begin()

	for i := 0; i < len(vals); i++ {
		pv := &vals[i]
		tx.Model(pv).Select("Value").Updates(pv)
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	// сразу обновим кеши, пока под локами
	r.cache = settingsCache{}

	return r.fetchNoLock()
}

func (r *Settings) Migrate() (err error) {
	return r.db.AutoMigrate(&model.Setting{})
}

func (r *Settings) fetchNoLock() (err error) {

	var list model.SettingsList

	if err = r.db.Find(&list).Error; err != nil {
		return err
	}

	n := uint(len(list))

	if n == 0 {
		// TODO return errSettingFetchEmpty
		return nil
	}

	// build indexes
	index := makeEntityIndex(int(n))
	names := make(settingsNamesIndex, n)

	for i := uint(0); i < n; i++ {
		v := &list[i]
		index[v.Id.Raw()] = i
		names[v.Name] = i
	}

	// store indexes
	r.cache = settingsCache{
		list:  list,
		index: index,
		names: names,
	}

	return nil
}

func (r *Settings) Run() error {
	return r.fetchNoLock()
}

func (r *Settings) Stop() error {
	return nil
}

func NewSettings(db *gorm.DB) (r *Settings, err error) {

	r = &Settings{
		db: db,
	}

	/*
		if err = r.fetchNoLock(); err != nil {
			return nil, err
		}
	*/

	return r, nil
}
