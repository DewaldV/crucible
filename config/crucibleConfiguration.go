package config

import (
  "fmt"
)

type CrucibleConfiguration struct {
  HttpPort      int
  HttpsPort     int
  WorkerThreads int
  RootContext   string
  CoreLogger    *LoggerConfig
  DataSources   map[string]*DataSourceConfiguration
  Services      map[string]*ServiceConfiguration
}

func (c *CrucibleConfiguration) String() (s string) {
  s = fmt.Sprintf("HttpPort: %d\n", c.HttpPort)
  s += fmt.Sprintf("HttpsPort: %d\n", c.HttpsPort)
  s += fmt.Sprintf("RootContext: %s\n", c.RootContext)
  s += fmt.Sprintf("WorkerThreads: %d\n", c.WorkerThreads)
  for key, value := range c.DataSources {
    s += fmt.Sprintf("DataSourceName: %s\n", key)
    s += value.String()
  }
  for key, value := range c.Services {
    s += fmt.Sprintf("ServiceName: %s\n", key)
    s += value.String()
  }
  return
}


func newCrucibleConfiguration() (c *CrucibleConfiguration) {
  c = new(CrucibleConfiguration)
  c.HttpPort = 8787
  c.HttpsPort = 44443
  c.RootContext = "/"
  c.WorkerThreads = 1
  return
}