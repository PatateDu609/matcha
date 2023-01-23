package config

type Session struct {
	Provider   string `mapstructure:"provider"`
	CookieName string `mapstructure:"cookie-name"`
	Lifetime   int64  `mapstructure:"lifetime"`
}
