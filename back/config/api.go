package config

type API struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	FrontURL string `mapstructure:"front"`
}
