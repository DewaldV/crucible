package database

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type DataSourceConfigStruct struct {
	ServerName   string
	ServerPort   int
	DatabaseName string
}

func (d *DataSourceConfigStruct) PrintConfig() (s string) {
	s = fmt.Sprintf("\t> ServerName: %s\n", d.ServerName)
	s += fmt.Sprintf("\t> ServerPort: %d\n", d.ServerPort)
	s += fmt.Sprintf("\t> DatabaseName: %s\n", d.DatabaseName)
	return
}

type MgoConnection struct {
	Database   string
	Collection string
}

func (conn *MgoConnection) FindOne(m bson.M, d interface{}) {
	ExecuteWithCollection(conn.Database, conn.Collection, func(c *mgo.Collection) error { return c.Find(m).One(d) })
}

func (conn *MgoConnection) FindAll(m bson.M, d interface{}) {
	ExecuteWithCollection(conn.Database, conn.Collection, func(c *mgo.Collection) error { return c.Find(m).All(d) })
}

/*
func (conn *MgoConnection) Update(selector interface{}, d interface{}) {
	ExecuteWithCollection(conn.Database, conn.Collection, func(c *mgo.Collection) error { return c.Update(selector, d) })
}

func (conn *MgoConnection) UpdateAll(selector interface{}, d interface{}) {
	ExecuteWithCollection(conn.Database, conn.Collection, func(c *mgo.Collection) error { return c.UpdateAll(selector, d) })
}

func (conn *MgoConnection) Upsert(selector interface{}, d interface{}) {
	ExecuteWithCollection(conn.Database, conn.Collection, func(c *mgo.Collection) error { return c.Upsert(selector, d) })
}
*/
func (conn *MgoConnection) Insert(d interface{}) {
	ExecuteWithCollection(conn.Database, conn.Collection, func(c *mgo.Collection) error { return c.Insert(d) })
}

/*
func (conn *MgoConnection) Remove(d interface{}) {
	ExecuteWithCollection(conn.Database, conn.Collection, func(c *mgo.Collection) error { return c.Remove(d) })
}

func (conn *MgoConnection) RemoveAll(d interface{}) {
	ExecuteWithCollection(conn.Database, conn.Collection, func(c *mgo.Collection) error { return c.RemoveAll(d) })
}
*/

func ExecuteWithCollection(database, collection string, f func(*mgo.Collection) error) error {
	session := GetSession("Default")
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database).C(collection)

	return f(c)
}

var sessionPool map[string]*mgo.Session

func LoadSessions(dataSourceConfig map[string]*DataSourceConfigStruct) {
	sessionPool = make(map[string]*mgo.Session)
	for key, source := range dataSourceConfig {
		s, _ := mgo.Dial(source.ServerName)
		sessionPool[key] = s
	}
}

func GetSession(dataSource string) *mgo.Session {
	s := sessionPool[dataSource]
	return s.Clone()
}
