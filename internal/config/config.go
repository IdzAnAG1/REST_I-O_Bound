package config

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"rest_io_bound/internal/variables"
)

type Config struct {
	// http configurations
	HTTPServerConfig struct {
		Port string
	}
}

func LoadConfig() (*Config, error) {
	if !isConfigDirectoryExists(variables.CONFIG_DIRECTORY) {
		if err := createConfigDirectory(
			variables.CONFIG_DIRECTORY); err != nil {
			return nil, err
		}
		if err := createConfigFile(
			variables.CONFIG_DIRECTORY, variables.CONFIG_FILE_NAME+variables.CONFIG_FILE_TYPE); err != nil {
			return nil, err
		}
	}
	cfg, err := readConfig(variables.CONFIG_DIRECTORY)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
func createConfigDirectory(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}
	return nil
}
func createConfigFile(path, filename string) error {
	fullPath := filepath.Join(path, filename)
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()
	err = os.WriteFile(fullPath, []byte(variables.DEFAULT_VALUE_FOR_CONFIG), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
func isConfigDirectoryExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return fileInfo.IsDir()
}

func readConfig(path string) (*Config, error) {
	cfg := &Config{}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
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
