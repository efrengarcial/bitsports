package datasource

import (
	"bitsports/config"
	"bitsports/ent"
	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"net/url"
)

// New returns data source name
func New() string {
	sslMode := "require"
	if config.C.Database.Params.DisableTLS {
		sslMode = "disable"
	}

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", config.C.Database.Params.Timezone)

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(config.C.Database.User, config.C.Database.Password),
		Host:     config.C.Database.Addr,
		Path:     config.C.Database.DBName,
		RawQuery: q.Encode(),
	}

	return u.String()
}

// NewClient returns an orm client
func NewClient() (*ent.Client, error) {
	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	d := New()

	return ent.Open(dialect.Postgres, d, entOptions...)
}

