package constants

import "os"

var (
	ENV_DATABASE_URL = lookupEnv("DATABASE_URL", "postgres://example:example@example-db:5432/tracking?sslmode=disable")
)

func lookupEnv(envname string, defaultVal string) string {
	if v, ok := os.LookupEnv(envname); ok {
		return v
	}
	return defaultVal
}
