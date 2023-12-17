package entities

type Config struct {
	Components struct {
		Service  ConfigService  `yaml:"service"`
		Telegram ConfigTelegram `yaml:"telegram"`
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
