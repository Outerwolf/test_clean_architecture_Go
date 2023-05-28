package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type ConfigBuilder struct {
	basePath   string
	configType string
	viper      *viper.Viper
}

func NewViperBuilder(options ...viper.Option) ConfigBuilder {
	return ConfigBuilder{basePath: "configuration", viper: viper.NewWithOptions(options...)}
}

func (cb ConfigBuilder) WithConfigType(configType string) ConfigBuilder {
	cb.configType = configType
	return cb
}

func (cb ConfigBuilder) WithBasePath(basePath string) ConfigBuilder {
	cb.basePath = basePath
	return cb
}

func (cb ConfigBuilder) WithConfigFile(filePath string, optional bool) ConfigBuilder {

	viperFile := viper.New()

	viperFile.SetConfigType(cb.configType)
	viperFile.AddConfigPath(cb.basePath)
	viperFile.SetConfigName(filePath)
	err := viperFile.ReadInConfig()
	if err != nil {
		if optional {
			return cb
		}
		panic(err)
	}

	err = cb.viper.MergeConfigMap(viperFile.AllSettings())
	if err != nil {
		panic(err)
	}
	return cb
}

func (cb ConfigBuilder) Build() *viper.Viper {
	return cb.viper
}

func NewConfigBuilder() *viper.Viper {
	var environment string = "Local"

	return NewViperBuilder().
		WithConfigType("yaml").
		WithBasePath("configuration").
		WithConfigFile("config", true).
		WithConfigFile(fmt.Sprintf("config.%s", environment), true).
		Build()
}

func NewConfiguration(v *viper.Viper) *Configuration {
	c := &Configuration{}
	if err := v.Unmarshal(c); err != nil {
		panic(err)
	}
	return c
}
