package main

import (
	"flag"
	"github.com/pkg6/oamc/sls/config"
	"github.com/pkg6/oamc/sls/lang"
	"github.com/pkg6/oamc/sls/listener"
	"github.com/pkg6/oamc/sls/sender"
	"log"
	"os"
	"path"
)

var (
	homePath  string
	paths     []string
	cfg       *config.Config
	err       error
	inputPath string
)

func init() {
	_ = os.Setenv("SLS_LANG", "zh-CN")
	homePath, _ = os.UserHomeDir()
	paths = []string{
		path.Join(homePath, ".oamc", "sls.json"),
		path.Join("/etc", "oamc", "sls.json"),
	}
	lang.Load()
	flag.StringVar(&inputPath, "f", "config.json", "User-defined configuration files")
}

func main() {
	flag.Parse()
	paths = append(paths, inputPath)
	cfg, err = config.LoadConfig(paths)
	if err != nil {
		log.Fatal(err)
		return
	}
	sender.New(listener.New(cfg)).Run()
}
