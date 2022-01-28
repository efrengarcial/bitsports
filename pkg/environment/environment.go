package environment

import "os"

// Application Environment name
const (
	Test        = "test"
	Prod		= "prod"
	Mig			= "mig"
)


// IsTest returns APP_ENV in test mode
func IsTest() bool {
	return os.Getenv("APP_ENV") == Test
}

// IsProd returns APP_ENV in prod mode
func IsProd() bool {
	return os.Getenv("APP_ENV") == Prod
}


// IsMig returns APP_ENV in mig mode
func IsMig() bool {
	return os.Getenv("APP_ENV") == Mig
}
