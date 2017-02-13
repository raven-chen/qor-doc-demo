package config

import (
	"fmt"
	"os"
	"path"

	"github.com/jinzhu/configor"
	"github.com/theplant/appkit/server"
)

var Config = struct {
	Env string `env:"ENV" default:"development"`
	DB  struct {
		Name     string
		Host     string
		Port     string
		Adapter  string
		User     string
		Password string
	}
	SessionStoreKey string `yaml:"session_store_key"`

	HTTP server.Config
}{}

var Root string

var ENV = ""

func init() {
	Root = path.Join(os.Getenv("GOPATH"), "/src/github.com/raven-chen/qor_doc_demo")

	if os.Getenv("QDD_ENV") != "" {
		os.Setenv("CONFIGOR_ENV", os.Getenv("QDD_ENV"))
	}

	ENV = configor.ENV()
	fmt.Printf("ENV: %s\n", ENV)

	if err := configor.Load(&Config, path.Join(Root, "config/config.yml")); err != nil {
		panic(err)
	}

	fmt.Printf("database name: %s\n", Config.DB.Name)

}

func IsTestEnv() bool {
	return configor.ENV() == "test"
}

func IsStaging() bool {
	return ENV == "staging"
}

func IsDev() bool {
	return ENV == "development"
}
