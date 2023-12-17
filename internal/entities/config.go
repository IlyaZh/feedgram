package entities

type Config struct {
	Components struct {
		Telegram Telegram `yaml:"telegram"`
	} `yaml:"components"`
}

type Telegram struct {
	Token   string `yaml:"token"`
	Limit   *int   `yaml:"limit"`
	Timeout *int   `yaml:"timeout"`
}
