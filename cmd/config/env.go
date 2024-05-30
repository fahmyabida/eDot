package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func mustLoad(prefix string, spec interface{}) {
	err := envconfig.Process(prefix, spec)
	if err != nil {
		panic(err)
	}
}

// InitEnv initializes environment variables for local testing from .env file
func InitEnv(envFileName ...string) error {
	if IsDevEnv() {
		_, thisFilePath, _, _ := runtime.Caller(0)
		dotEnvPath := filepath.Join(filepath.Dir(thisFilePath), "../../")
		if envFileName != nil {
			dotEnvPath = filepath.Join(dotEnvPath, envFileName[0])
		} else {
			dotEnvPath = filepath.Join(dotEnvPath, ".env")
		}
	
		if err := godotenv.Load(dotEnvPath); err != nil {
			return fmt.Errorf("error loading %s file", dotEnvPath)
		}
	}

	return nil
}

// InitTestEnv wraps the generic InitEnv to accept a testing env file
func InitTestEnv() error {
	return InitEnv(".env.test")
}

func IsDevEnv() bool {
	appEnv := os.Getenv("APP_ENV")

	switch appEnv {
	case
		"",
		"dev",
		"local":
		return true
	}
	return false
}
