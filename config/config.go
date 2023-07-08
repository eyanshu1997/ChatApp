package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type RegistrationConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (c *RegistrationConfig) ReadConfigFile(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("unable to read open file %s", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(c)
	if err != nil {
		return fmt.Errorf("unable to decode config file %s", err)
	}
	return nil
}
