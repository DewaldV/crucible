package config

import "fmt"

type ServiceConfiguration struct {
  Location       string
  AllowedOrigins map[string]bool
  AllowedMethods map[string]bool
}

func (c *ServiceConfiguration) PrintConfig() (s string) {
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

func newServiceConfiguration() (s *ServiceConfiguration) {
  s = new(ServiceConfiguration)
  s.Location = "/"
  s.AllowedOrigins = make(map[string]bool)
  s.AllowedOrigins["localhost"] = true
  return
}