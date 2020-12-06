package filmsconf

import (
	"fmt"
	"os"
)

const DATABASE_SCHEME = "postgres"

func GetDATABASE_PORT() (port string) {
	return ":5432"
}

func GetDATABASE_NAME() (db string) {
	return "films_dev"
}

func GetDATABASE_USER() (user string) {
	return os.Getenv("POSTGRES_USER")
}

func GetDATABASE_HOST() (host string) {
	return "films_db"
}

func GetDATABASE_PASSWORD() (pwd string) {
	return os.Getenv("POSTGRES_PASSWORD")
}

func GetDATABASE_SRC() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s%s/%s?sslmode=%s",
		DATABASE_SCHEME, GetDATABASE_USER(), GetDATABASE_PASSWORD(), GetDATABASE_HOST(), GetDATABASE_PORT(),
		GetDATABASE_NAME(), GetDATABASE_SSLMODE(),
	)
}

func GetDATABASE_SSLMODE() (sslmode string) {
	return "disable" // just in dev
}
