package main

import (
	"github.com/kooksee/html_meta/internal/config"
	"github.com/kooksee/html_meta/internal/kts"
	"github.com/kooksee/html_meta/internal/router"
	"github.com/kooksee/html_meta/internal/utils"
)

func main() {
	cfg := config.DefaultConfig()
	cfg.Init()
	kts.Init()
	utils.MustNotError(router.App().Run())
}
