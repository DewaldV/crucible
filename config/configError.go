package config

import (
  "fmt"
)

type ConfigError struct {
  ConfigFile string
  What       string
  Err        string
}

func (e *ConfigError) Error() string {
  return fmt.Sprintf("Configuration error in %s > %s > %s\n", e.ConfigFile, e.What, e.Err)
}