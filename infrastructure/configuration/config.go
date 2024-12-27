package configuration

import "github.com/spf13/viper"

type MarkitosGolangServiceBoilerplateConfig struct {
	DsnDatabase string `mapstructure:"APP_BBDD_DSN"`
	AppAddress  string `mapstructure:"APP_ADDRESS"`
}

func LoadConfiguration(
	configFilesPath string) (config MarkitosGolangServiceBoilerplateConfig, err error) {

	viper.AddConfigPath(configFilesPath)

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)

	return config, err
}
