package environment

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var configManager *viper.Viper

var (
	ErrValidation    = errors.New("validation error")
	ErrConfigManager = errors.New("config manager error")
)

type EnvLoaderParams struct {
	File string
	Path string
	Type string
}

type EnvLoader interface {
	Load() error
}

type envLoader struct {
	params *EnvLoaderParams
}

func NewEnvLoader(params *EnvLoaderParams) error {
	if params == nil {
		params = &EnvLoaderParams{
			File: ".env",
			Path: ".",
			Type: "env",
		}
	}

	env := &envLoader{
		params: params,
	}

	err := env.Load()
	if err != nil {
		return err
	}

	return nil
}

func (l *envLoader) Load() error {
	err := l.newConfigManager()
	if err != nil {
		return fmt.Errorf("%w: %s", ErrConfigManager, err)
	}

	configFileExists := l.envFileExists(path.Join(l.params.Path, l.params.File))
	if configFileExists {
		err = l.validate()
		if err != nil {
			return fmt.Errorf("%w: %s", ErrValidation, err)
		}
	}

	return nil
}

func (l *envLoader) validate() error {
	err := validator.New().Struct(&envSchema)
	if err != nil {
		return err
	}

	return nil
}

func (l *envLoader) newConfigManager() error {
	localConfigManager := viper.New()
	localConfigManager.SetConfigName(fmt.Sprintf("%s_environment", path.Base(l.params.Path)))

	configFilePath := path.Join(l.params.Path, l.params.File)
	if l.envFileExists(configFilePath) {
		localConfigManager.AddConfigPath(l.params.Path)
		localConfigManager.SetConfigType(l.params.Type)
		localConfigManager.SetConfigFile(configFilePath)

		err := localConfigManager.ReadInConfig()
		if err != nil {
			return err
		}
	}

	localConfigManager.AutomaticEnv()

	err := localConfigManager.Unmarshal(&envSchema)
	if err != nil {
		return err
	}

	configManager = localConfigManager

	return nil
}

func (l *envLoader) envFileExists(envPath string) bool {
	_, err := os.Stat(envPath)
	return !os.IsNotExist(err)
}
