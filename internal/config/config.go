package config

import (
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"rest_io_bound/internal/variables"
)

type Config struct {
	HTTP HTTP `mapstructure:"http"`
}
type HTTP struct {
	Port string `mapstructure:"port"`
}

func LoadConfig() (*Config, error) {
	dirPath, err := pathToConfigDir()
	if err != nil {
		return nil, err
	}
	filePath := pathToConfigFile(dirPath)
	if !IsConfigDirectoryExists(dirPath) {
		if err := createConfigDirectory(dirPath); err != nil {
			return nil, err
		}
		if err := createConfigFile(filePath); err != nil {
			return nil, err
		}
	}
	cfg, err := readConfig(dirPath)

	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func createConfigDirectory(pathToDir string) error {
	return os.MkdirAll(pathToDir, os.ModePerm)
}

func createConfigFile(pathToFile string) error {
	cfg := Config{
		HTTP: HTTP{
			Port: variables.DEFAULT_HTTP_PORT,
		},
	}

	data, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(pathToFile, data, os.ModePerm)
}

func IsConfigDirectoryExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return fileInfo.IsDir()
}

func readConfig(path string) (*Config, error) {
	cfg := &Config{}
	viper.SetConfigName(variables.CONFIG_FILE_NAME)
	viper.SetConfigType(variables.CONFIG_FILE_TYPE)
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func pathToConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, variables.CONFIG_DIRECTORY), nil
}

func pathToConfigFile(pathToDir string) string {
	return filepath.Join(pathToDir, variables.CONFIG_FILE_NAME+"."+variables.CONFIG_FILE_TYPE)
}
