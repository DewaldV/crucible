package config

import (
  "fmt"
)

type DataSourceConfiguration struct {
  ServerName   string
  ServerPort   int
  DatabaseName string
}

func (d *DataSourceConfiguration) String() (s string) {
  s = fmt.Sprintf("\t> ServerName: %s\n", d.ServerName)
  s += fmt.Sprintf("\t> ServerPort: %d\n", d.ServerPort)
  s += fmt.Sprintf("\t> DatabaseName: %s\n", d.DatabaseName)
  return
}
