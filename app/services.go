/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package app

import (
	"github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/config"
	"gorm.io/gorm"
)

type services struct {
	// service *service.Service // or proto.ServiceService
}

/*
//go:nosplit
func (s *services) Service() *service.Service { // or proto.ServiceService
	return s.service
}
*/

func (s *services) init(cfg *config.Config, db *gorm.DB) (err error) {

	/* TODO
	s.service, err = service.NewService(db)

	if err != nil {
		return fmt.Errorf("Service service init error: %w", err)
	}

	*/

	//

	// s.serviceWOErr = service.NewServiceWOError()

	return nil
}

func (s *services) run() (err error) {

	// ...

	return nil
}

func (s *services) stop() (err error) {

	/* TODO
	if err = s.service.Stop(); err != nil {
		return fmt.Errorf("Service service stop error: %w", err)
	}
	*/

	return nil
}

// ------------------------------------------------------------

//go:nosplit
func (app *App) initServices() (err error) {
	return app.services.init(app.cfg, app.db)
}
