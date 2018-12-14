package main

import (
	"github.com/kooksee/html_meta/internal/config"
	"github.com/kooksee/html_meta/internal/kts"
	"github.com/kooksee/html_meta/internal/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	cfg := config.DefaultConfig()
	cfg.Init()
	kts.Init()

	utils.MustNotError(filepath.Walk("patterns", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		dt, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		hp := &kts.HtmlPattern{
			Name:    strings.Split(info.Name(), ".")[0],
			Pattern: string(dt),
		}

		if hp.Exist() {
			return hp.Update()
		}
		return hp.Save()
	}))
}
