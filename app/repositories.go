/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package app

import (
	"fmt"

	"github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/config"
	"github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/proto"
	"github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/repository"
	"gorm.io/gorm"
)

type repos struct {
	settings *repository.Settings
}

//go:nosplit
func (r *repos) Settings() proto.RepositorySettings {
	return r.settings
}

func (r *repos) init(cfg *config.Config, db *gorm.DB) (err error) {

	r.settings, err = repository.NewSettings(db)

	if err != nil {
		return fmt.Errorf("Settings repo init error: %w", err)
	}

	//

	// s.repoWOErr = repository.NewRepoWOError()

	// if cfg.migrate

	if err = r.settings.Migrate(); err != nil {
		return fmt.Errorf("Settings migrate error: %w", err)
	}

	return nil
}

func (r *repos) run() (err error) {

	if err = r.settings.Run(); err != nil {
		return err
	}

	return nil
}

func (r *repos) stop() (err error) {

	/* TODO
	if err = s.service.Stop(); err != nil {
		return fmt.Errorf("Service service stop error: %w", err)
	}
	*/

	return nil
}

// ------------------------------------------------------------

//go:nosplit
func (app *App) initRepos() (err error) {
	return app.repos.init(app.cfg, app.db)
}
