package config

import "github.com/spf13/viper"

type Config struct {
	DBPort         string `mapstructure:"DB_PORT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBName         string `mapstructure:"DB_NAME"`
	DBUsername     string `mapstructure:"DB_USERNAME"`
	DBPass         string `mapstructure:"DB_PASSWORD"`
	DBDialect      string `mapstructure:"DB_DIALECT"`
	SUPERUSERPASS  string `mapstructure:"SUPERUSERPASS"`
	Env            string `mapstructure:"ENV"`
	Port           string `mapstructure:"PORT"`
	AppName        string `mapstructure:"APP_NAME"`
	LogPath        string `mapstructure:"LOG_PATH"`
	CloudUser      string `mapstructure:"CLOUD_NAME"`
	CloudApiKey    string `mapstructure:"CLOUD_API_KEY"`
	CloudApiSecret string `mapstructure:"CLOUD_API_SECRET"`
	CloudUrl       string `mapstructure:"CLOUD_URL"`
	MaxUploadSize  int64  `mapstructure:"MAX_IMAGE_SIZE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
