package entities

// todo: mode secret info to env
type Config struct {
	Components struct {
		Service  ConfigService  `yaml:"service"`
		Telegram ConfigTelegram `yaml:"telegram"`
		Postgres ConfigPostgres `yaml:"postgres"`
	} `yaml:"components"`
}

type ConfigTelegram struct {
	Token   string `yaml:"token"`
	Limit   *int   `yaml:"limit"`
	Timeout *int   `yaml:"timeout"`
}

type ConfigService struct {
	Port int `yaml:"port"`
}

type ConfigPostgres struct {
	Host               *string `yaml:"host"`
	User               string  `yaml:"user"`
	Password           string  `yaml:"password"`
	Port               *int    `yaml:"port"`
	Database           string  `yaml:"database"`
	SslMode            *string `yaml:"sslmode"`
	MaxOpenConnections *int    `yaml:"max_open_connections"`
	MaxIdleConnections *int    `yaml:"max_idle_connections"`
	Limit              *int    `yaml:"limit"`
}
