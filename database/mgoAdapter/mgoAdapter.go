package mgoAdapter

import (
	"labix.org/v2/mgo"
)

type Connection struct {
	Database   string
	Collection string
}

func (conn *Connection) Connect(serverName string) (interface{}, error) {
	return mgo.Dial(serverName)
}

func (conn *Connection) FindOne(selector interface{}, d interface{}) error {
	_, err := ExecuteWithCollection(conn.Database, conn.Collection, func(c *mgo.Collection) (*mgo.ChangeInfo, error) { return nil, c.Find(selector).One(d) })
	return err
}

func (conn *Connection) FindAll(selector interface{}, d interface{}) error {
	_, err := ExecuteWithCollection(conn.Database, conn.Collection, func(c *mgo.Collection) (*mgo.ChangeInfo, error) { return nil, c.Find(selector).All(d) })
	return err
}

func (conn *Connection) Update(selector interface{}, d interface{}) error {
	_, err := ExecuteWithCollection(conn.Database, conn.Collection, func(c *mgo.Collection) (*mgo.ChangeInfo, error) { return nil, c.Update(selector, d) })
	return err
}

func (conn *Connection) UpdateAll(selector interface{}, d interface{}) (int, error) {
	changeInfo, err := ExecuteWithCollection(conn.Database, conn.Collection, func(c *mgo.Collection) (*mgo.ChangeInfo, error) { return c.UpdateAll(selector, d) })
	return changeInfo.Updated, err
}

func (conn *Connection) Upsert(selector interface{}, d interface{}) (interface{}, error) {
	changeInfo, err := ExecuteWithCollection(conn.Database, conn.Collection, func(c *mgo.Collection) (*mgo.ChangeInfo, error) { return c.Upsert(selector, d) })
	return changeInfo.UpsertedId, err
}

func (conn *Connection) Insert(d interface{}) error {
	_, err := ExecuteWithCollection(conn.Database, conn.Collection, func(c *mgo.Collection) (*mgo.ChangeInfo, error) { return nil, c.Insert(d) })
	return err
}

func (conn *Connection) Remove(d interface{}) error {
	_, err := ExecuteWithCollection(conn.Database, conn.Collection, func(c *mgo.Collection) (*mgo.ChangeInfo, error) { return nil, c.Remove(d) })
	return err
}

func (conn *Connection) RemoveAll(d interface{}) (int, error) {
	changeInfo, err := ExecuteWithCollection(conn.Database, conn.Collection, func(c *mgo.Collection) (*mgo.ChangeInfo, error) { return c.RemoveAll(d) })
	return changeInfo.Removed, err
}

func ExecuteWithCollection(database, collection string, f func(*mgo.Collection) (*mgo.ChangeInfo, error)) (*mgo.ChangeInfo, error) {
	session := GetSession(database)
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database).C(collection)
	changeInfo, err := f(c)
	return changeInfo, err
}

func GetSession(dataSource string) *mgo.Session {
	s := sessionPool[dataSource]
	return s.Clone()
}
