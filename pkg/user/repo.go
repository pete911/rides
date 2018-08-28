package user

import (
	"fmt"
	"github.com/pete911/rides/pkg/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ErrNotFound struct {
	err string
}

func (e *ErrNotFound) Error() string {
	return e.err
}

func Add(user *User) error {

	b, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("add user: bcrypt generate from password: %v", err)
	}

	user.Password = b
	user.Created = time.Now()
	user.Id = bson.NewObjectId().Hex()

	s := mongo.Session.Copy()
	defer s.Close()

	_, err = s.DB(mongo.DbName).C(mongo.UserCollectionId).UpsertId(user.Id, user)
	return err
}

func Get(email string) (User, error) {

	s := mongo.Session.Copy()
	defer s.Close()

	var user User
	if err := s.DB(mongo.DbName).C(mongo.UserCollectionId).Find(bson.M{"email": email}).One(&user); err != nil {
		if err == mgo.ErrNotFound {
			return User{}, &ErrNotFound{err: fmt.Sprintf("user %s not found", email)}
		}
		return User{}, err
	}
	return user, nil
}

func FindAll() ([]User, error) {

	s := mongo.Session.Copy()
	defer s.Close()

	var users []User
	err := s.DB(mongo.DbName).C(mongo.UserCollectionId).Find(nil).All(&users)

	return users, err
}

func Remove(email string) error {

	s := mongo.Session.Copy()
	defer s.Close()

	return s.DB(mongo.DbName).C(mongo.UserCollectionId).Remove(bson.M{"email": email})
}
