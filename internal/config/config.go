package config

// Config is a struct to hold config value froom yaml file
// it use mapstructure tag to map each key because viper use mapstructure and not json or yaml to unnmarshal
type Config struct {
	App App `mapstructure:"app" yaml:"app"`
}

type App struct {
	Name         string `mapstructure:"name" yaml:"name"`
	Port         string `mapstructure:"port" yaml:"port"`
	ReadTimeout  int    `mapstructure:"read_timeout" yaml:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout" yaml:"write_timeout"`
	ReleaseMode  string `mapstructure:"release_mode" yaml:"release_mode"`
}
