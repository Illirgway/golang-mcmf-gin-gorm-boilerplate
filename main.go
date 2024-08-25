/**
 *
 * Copyright (c) 2024 Illirgway
 *
 */

package main

import (
	"log"

	"github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/app"
	"github.com/Illirgway/golang-mcmf-gin-gorm-boilerplate/config"
)

// WIN:
// go build -v -o bin\boilerplate.exe .
// go tool objdump -S -s "boilerplate" bin\boilerplate.exe > bin\boilerplate.disasm
//
// bin\boilerplate.exe -c dist\config.json
//
// go build -v -ldflags "-s -w" -o bin\boilerplate.exe .
//

func main() {

	cfg, err := config.New()

	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	app := app.New(cfg)

	if err = app.Setup(); err != nil {
		log.Fatalf("App setup error: %v", err)
	}

	if err = registerControllers(app); err != nil {
		log.Fatalf("Controller registration error: %v", err)
	}

	if err = app.Run(); err != nil {
		log.Fatalf("App run error: %v", err)
		return
	}

	log.Println("exit ok...")
}
