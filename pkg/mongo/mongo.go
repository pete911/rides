package mongo

import (
	"gopkg.in/mgo.v2"
	"log"
)

const (
	DbName           = "ride"
	UserCollectionId = "user"
)

// Session needs to be initialized 'InitSession(url)' before it can be used
var Session *mgo.Session

func Health() error {

	s := Session.Copy()
	defer s.Close()
	return s.Ping()
}

// Call once, preferably in main method before application starts
func InitSession(url string) {

	if Session != nil {
		log.Print("mongo session is already initialized")
		return
	}

	s, err := mgo.Dial(url)
	if err != nil {
		log.Fatalf("Unable to connect to mongo on url=%s: %v", url, err)
	}
	Session = s

	// ensure indexes
	ensureIndex(UserCollectionId, mgo.Index{Name: "userEmail", Key: []string{"email"}, Unique: true})
}

func ensureIndex(collection string, index mgo.Index) {

	s := Session.Copy()
	defer s.Close()

	if err := s.DB(DbName).C(collection).EnsureIndex(index); err != nil {
		log.Fatalf("Error ensuring index: '%+v', collection: '%s', error: '%v'", index.Key, collection, err)
	}
}
