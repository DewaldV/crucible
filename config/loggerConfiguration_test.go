package config

import (
  "testing"
  "github.com/DewaldV/crucible/logging"
)

func TestNewLoggerConfig(t *testing.T) {
  l := newLoggerConfig()

  if(l.Level != logging.Info || l.FileName != "log") {
    t.Errorf("Logger not created correctly. Expected Level: %s, FileName: log | Found Level: %s, FileName: %s", logging.Info, l.Level, l.FileName)
  }
}