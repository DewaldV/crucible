package config

import (
  "github.com/DewaldV/crucible/logging"
)

type LoggerConfig struct {
  Level    logging.LogLevel
  FileName string
}

func newLoggerConfig() (c *logging.LoggerConfig) {
  c = new(logging.LoggerConfig)
  c.FileName = "log"
  c.Level = logging.Info
  return
}