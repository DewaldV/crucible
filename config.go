package crucible

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/DewaldV/crucible/database"
	"github.com/DewaldV/crucible/logging"
	"io/ioutil"
	"os"
)

type ConfigPrinter interface {
	PrintConfig() string
}

type CoreConfigStruct struct {
	HttpPort      int
	HttpsPort     int
	WorkerThreads int
	RootContext   string
	DataSources   map[string]*database.DataSourceConfigStruct
	Services      map[string]*ServiceConfigStruct
}

type ServiceConfigStruct struct {
	Location       string
	AllowedOrigins map[string]bool
	AllowedMethods map[string]bool
}

type LoggerConfig struct {
	Level    logging.LogLevel
	FileName string
}

var Conf *CoreConfigStruct

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

	Conf = newCoreConfigStruct()

	err = json.Unmarshal(contents, Conf)
	if err != nil {
		return &ConfigError{path, "Error parsing config file", err.Error()}
	}

	fmt.Println("Loaded configuration:")
	fmt.Println(Conf.PrintConfig())

	database.LoadSessions(Conf.DataSources)

	return nil
}

func newCoreConfigStruct() (c *CoreConfigStruct) {
	c = new(CoreConfigStruct)
	c.HttpPort = 8787
	c.HttpsPort = 44443
	c.RootContext = "/"
	c.WorkerThreads = 1
	return
}

func newSiteConfigStruct() (s *ServiceConfigStruct) {
	s = new(ServiceConfigStruct)
	s.Location = "/"
	s.AllowedOrigins = make(map[string]bool)
	s.AllowedOrigins["localhost"] = true
	return
}

func newLoggerConfig() (c *logging.LoggerConfig) {
	c = new(logging.LoggerConfig)
	c.FileName = "log"
	c.Level = logging.Info
	return
}

func (c *CoreConfigStruct) PrintConfig() (s string) {
	s = fmt.Sprintf("HttpPort: %d\n", c.HttpPort)
	s += fmt.Sprintf("HttpsPort: %d\n", c.HttpsPort)
	s += fmt.Sprintf("RootContext: %s\n", c.RootContext)
	s += fmt.Sprintf("WorkerThreads: %d\n", c.WorkerThreads)
	for key, value := range c.DataSources {
		s += fmt.Sprintf("DataSourceName: %s\n", key)
		s += value.PrintConfig()
	}
	for key, value := range c.Services {
		s += fmt.Sprintf("ServiceName: %s\n", key)
		s += value.PrintConfig()
	}
	return
}

func (c *ServiceConfigStruct) PrintConfig() (s string) {
	s = fmt.Sprintf("\t> Location: %s\n", c.Location)
	var allowedMethods string
	for key := range c.AllowedMethods {
		allowedMethods += fmt.Sprintf("%s,", key)
	}
	s += fmt.Sprintf("\t> AllowedMethods: %s\n", allowedMethods[:len(allowedMethods)-1])

	var allowedOrigins string
	for key := range c.AllowedOrigins {
		allowedOrigins += fmt.Sprintf("%s,", key)
	}
	s += fmt.Sprintf("\t> AllowedOrigins: %s\n", allowedOrigins[:len(allowedOrigins)-1])
	return
}

type ConfigError struct {
	ConfigFile string
	What       string
	Err        string
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("Configuration error in %s > %s > %s\n", e.ConfigFile, e.What, e.Err)
}
