package config

type DatabaseEnv struct {
	Hostname string `envconfig:"HOSTNAME"`
	Username string `envconfig:"USERNAME"`
	Password string `envconfig:"PASSWORD"`
	Port     string `envconfig:"PORT"`
	DBName   string `envconfig:"DBNAME"`
}
