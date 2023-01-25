package config

type API struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	InternalPort int    `mapstructure:"internal-port"`
	FrontURL     string `mapstructure:"front"`
}
