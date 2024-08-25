/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (app *App) initAppDb() (err error) {

	logCfg := logger.Config{
		SlowThreshold:        200 * time.Millisecond,
		Colorful:             true,
		ParameterizedQueries: false,
		LogLevel:             logger.Warn,
	}

	if app.Config().Debug {
		logCfg.LogLevel = logger.Info
	}

	lgr := logger.New(
		log.New(os.Stderr, "\r\n", log.LstdFlags),
		logCfg)

	cfg := gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 lgr,
	}

	rdbms := mysql.Open(app.cfg.GetDbDSN())

	if app.db, err = gorm.Open(rdbms, &cfg); err != nil {
		return fmt.Errorf("GORM.Open error: %w", err)
	}

	return nil
}
