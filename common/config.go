package common

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type ProviderConfig struct {
	Name    string                 `yaml:"name"`
	Options map[string]interface{} `yaml:"options"`
}

type Config struct {
	AuthProvider ProviderConfig `yaml:"authProvider"`
	RoleProvider ProviderConfig `yaml:"roleProvider"`
}

func ConfigFactory(filepath string) Config {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(fmt.Sprintf("%s not found.", filepath))
	}
	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(fmt.Sprintf("wrong %s found", filepath), err)
	}
	return config
}
