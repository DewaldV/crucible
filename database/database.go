package database

import (
	"fmt"
	"github.com/DewaldV/crucible/database/mgoAdapter"
)

type DBConnection interface {
	Connect(serverName string) (interface{}, error)
}

type DBSession interface {
	Clone() DBSession
}

type DataSource interface {
	DBConnection
	FindOne(selector, d interface{}) error
	FindAll(selector, d interface{}) error
	Update(selector, d interface{}) error
	UpdateAll(selector, d interface{}) (int, error)
	Upsert(selector, d interface{}) (interface{}, error)
	Insert(d interface{}) error
	Remove(d interface{}) error
	RemoveAll(d interface{}) (int, error)
}

type DataSourceConfigStruct struct {
	Driver       string
	ServerName   string
	ServerPort   int
	DatabaseName string
	Username     string
	Password     string
}

func (d *DataSourceConfigStruct) PrintConfig() (s string) {
	s = fmt.Sprintf("\t> ServerName: %s\n", d.ServerName)
	s += fmt.Sprintf("\t> ServerPort: %d\n", d.ServerPort)
	s += fmt.Sprintf("\t> DatabaseName: %s\n", d.DatabaseName)
	return
}

var sessionPool map[string]*DBSession

func GetMgoConnection(dConf *DataSourceConfigStruct) DataSource {
	conn := &mgoAdapter.Connection{dConf.DatabaseName, ""}
	return conn
}

func LoadSessions(dataSourceConfig map[string]*DataSourceConfigStruct) {
	sessionPool = make(map[string]*DBSession)
	for key, source := range dataSourceConfig {
		var s *DBSession
		if source.Driver == "mgo" {
			conn := &mgoAdapter.Connection{"test", "leads"}
			s = conn.Connect(source.ServerName)
		}

		sessionPool[key] = s
	}
}
