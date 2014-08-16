package config

import (
  "testing"
  "strings"
  "github.com/DewaldV/crucible/crutest"
)

func TestNewServiceConfiguration(t *testing.T) {
  s := newServiceConfiguration()

  crutest.Assert(t, crutest.EqualsAll( map[interface{}]interface{} { s.Location:"/", s.AllowedOrigins["localhost"]:true, s.AllowedMethods["GET"]:true } ))
}

func TestServiceConfigurationString(t *testing.T) {
  s := newServiceConfiguration()

  actualString := s.String()

  if(!strings.Contains(actualString, "Location: /") || !strings.Contains(actualString, "AllowedOrigins: localhost") || !strings.Contains(actualString, "AllowedMethods: GET")) {
    t.Errorf("Failed to print ServiceConfiguration")
  }
}