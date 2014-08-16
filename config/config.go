package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func init() {
	LoadDefaultConfig()
}

var DefaultConfigPaths = [...]string{
	"/etc/crucible/",
	"/usr/local/etc/crucible/",
	"~/",
	"~/Dev/",
	"./"}

const (
	DefaultConfigFileName = "crucible.conf"
)

type ConfigPrinter interface {
	PrintConfig() string
}

var Conf *CrucibleConfiguration

func LoadDefaultConfig() error {
	var configPath string
	for _, path := range DefaultConfigPaths {
		configPath = fmt.Sprintf("%s%s", path, DefaultConfigFileName)
		_, err := os.Open(configPath)
		if err == nil {
			return LoadConfig(configPath)
		}
	}

	return &ConfigError{"", "No default config file could be found", ""}
}

func LoadConfig(path string) error {
	configFile, err := os.Open(path)
	if err != nil {
		return &ConfigError{path, "Could not open config file for reading.", err.Error()}
	}
	defer configFile.Close()

	reader := bufio.NewReader(configFile)
	contents, err := ioutil.ReadAll(reader)
	if err != nil {
		return &ConfigError{path, "Error reading config file", err.Error()}
	}

	Conf = newCrucibleConfiguration()

	err = json.Unmarshal(contents, Conf)
	if err != nil {
		return &ConfigError{path, "Error parsing config file", err.Error()}
	}

	fmt.Println("Loaded configuration:", path)
	fmt.Println(Conf.PrintConfig())

	return nil
}
