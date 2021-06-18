package conf

import "os"

// PG Config
type PGConfig struct {
	HOST string
	PORT string
	USER string
	PWD  string
	DB   string
}

// redis Config
type RDConfig struct {
	HOST string
	PORT string
	DB   string
	PWD  string
}

func GetDatabaseConfig() (PGConfig, RDConfig) {
	var host, port, user, password, database string

	if host = os.Getenv("DATABASE_HOST"); len(host) == 0 {
		host = "10.10.10.10"
	}
	if port = os.Getenv("DATABASE_PORT"); len(port) == 0 {
		port = "31000"
	}
	if user = os.Getenv("DATABASE_USER"); len(user) == 0 {
		user = "postgres"
	}
	if password = os.Getenv("DATABASE_PASSWORD"); len(password) == 0 {
		password = "X5ax5a62ang"
	}
	if database = os.Getenv("DATABASE_DB"); len(database) == 0 {
		database = "erp"
	}

	return PGConfig{
			HOST: host,
			PORT: port,
			USER: user,
			PWD:  password,
			DB:   database,
		}, RDConfig{
			HOST: "",
			PORT: "",
			DB:   "",
			PWD:  "",
		}
}
