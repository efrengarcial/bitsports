package testutil

import (
	"bitsports/config"
	"bitsports/pkg/environment"
)

// ReadConfig reads config file for test.
func ReadConfig() {
	config.ReadConfig(config.ReadConfigOption{
		AppEnv: environment.Test,
	})
}
