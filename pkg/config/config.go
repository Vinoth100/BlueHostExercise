package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ServerPort string `yaml:"server_port"`
}

func LoadConfig(filename string, config *Config) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "Error reading config file")
	}
	return yaml.Unmarshal(buf, config)
}
