package config

import (
	"bitsports/pkg/environment"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		User     string
		Password string
		Addr     string
		DBName   string
		Params   struct {
			DisableTLS bool
			Timezone   string
		}
	}
	Server struct {
		Address string
	}
	UserServer struct {
		Address string
	}
	Debug bool
}

// C is config variable
var C config

// ReadConfigOption is a config option
type ReadConfigOption struct {
	AppEnv string
}

// ReadConfig configures config file
func ReadConfig(option ReadConfigOption) {
	Config := &C

	println("os.Getenv(APP_ENV): ", os.Getenv("APP_ENV"))

	if environment.IsProd() {
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config.prod")
	} else if environment.IsMig() {
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config.mig")
	} else if environment.IsTest() || (option.AppEnv == environment.Test) {
		fmt.Println(rootDir())
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config.test")
	} else {
		fmt.Println(rootDir())
		viper.AddConfigPath(filepath.Join(rootDir(), "config"))
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
