package infrastructure

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	User     string
	Password string
	Net      string
	Host     string
	Port     int
	DBName   string
}
