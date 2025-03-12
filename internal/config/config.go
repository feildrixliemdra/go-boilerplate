package config

// Config is a struct to hold config value from yaml file
// it use mapstructure tag to map each key because viper use mapstructure and not json or yaml to unnmarshal
type Config struct {
	App     App     `mapstructure:"app" yaml:"app"`
	Postgre Postgre `mapstructure:"postgre" yaml:"postgre"`
	MongoDB MongoDB `mapstructure:"mongodb" yaml:"mongodb"`
	JWT     JWT     `mapstructure:"jwt" yaml:"jwt"`
}

type App struct {
	Name         string `mapstructure:"name" yaml:"name"`
	Port         string `mapstructure:"port" yaml:"port"`
	ReadTimeout  int    `mapstructure:"read_timeout" yaml:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout" yaml:"write_timeout"`
	ReleaseMode  string `mapstructure:"release_mode" yaml:"release_mode"`
	SecretKey    string `mapstructure:"secret_key" yaml:"secret_key"`
	Env          string `mapstructure:"env" yaml:"env"`
}

type Postgre struct {
	IsEnabled   bool   `mapstructure:"is_enabled" yaml:"is_enabled" `
	URL         string `mapstructure:"url" yaml:"url" `
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
}

type MongoDB struct {
	IsEnabled bool   `mapstructure:"is_enabled" yaml:"is_enabled" `
	URL       string `mapstructure:"url" yaml:"url" `
}

type JWT struct {
	SecretKey string `mapstructure:"secret_key"`
}
