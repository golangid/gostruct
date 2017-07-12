package config

import (
	"strings"

	"github.com/spf13/viper"
)

type viperConfig struct {
}

func (v *viperConfig) init() {
	viper.SetEnvPrefix(`gostruct`)
	viper.AutomaticEnv()

	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

}

func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (v *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (v *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

func NewViperConfig() *viperConfig {
	v := &viperConfig{}
	v.init()
	return v
}
