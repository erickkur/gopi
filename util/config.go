package util

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func GetConfigInstance(filename string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(filename)

	switch os.Getenv("ENV") {
	case "staging":
		v.AddConfigPath(fmt.Sprintf("%s/../config/staging/", basepath))
	case "production":
		v.AddConfigPath(fmt.Sprintf("%s/../config/production/", basepath))
	default:
		v.AddConfigPath(fmt.Sprintf("%s/../config/local/", basepath))
	}

	v.AddConfigPath(fmt.Sprintf("%s/../config/", basepath))

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	return v
}
