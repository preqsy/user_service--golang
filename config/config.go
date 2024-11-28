package config

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

var ServiceName = "user_service"

type Secrets struct {
	User     string `json:"USER" envconfig:"USER"`
	DbName   string `json:"DBNAME" envconfig:"DBNAME"`
	Password string `json:"PASSWORD" envconfig:"PASSWORD"`
	Host     string `json:"HOST" envconfig:"HOST"`
	Port     string `json:"PORT" envconfig:"PORT"`
}

var ss Secrets

func init() {
	importPath := fmt.Sprintf("%s/config", strings.ReplaceAll(ServiceName, "-", "."))
	p, err := build.Default.Import(importPath, "", build.FindOnly)
	if err == nil {
		env := filepath.Join(p.Dir, "../.env")
		_ = godotenv.Load(env)

	}

	ss = Secrets{}

	ss.User = os.Getenv("USER")
	ss.DbName = os.Getenv("DBNAME")
	ss.Password = os.Getenv("PASSWORD")
	ss.Host = os.Getenv("HOST")
	ss.Port = os.Getenv("PORT")
}

func GetSecrets() Secrets {
	return ss
}
