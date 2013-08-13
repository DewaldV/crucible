package crucible

import (
	"labix.org/v2/mgo"
)

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
	s, ok := sessionPool[dataSource]
	if !ok {
		s, _ := mgo.Dial(Conf.DataSources[dataSource].ServerName)
		sessionPool[dataSource] = s
	}
	return s.Clone()
}
