package user

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id       string    `bson:"_id"`
	Email    string    `bson:"email"`
	Password []byte    `bson:"password"`
	Roles    []string  `bson:"roles"`
	Created  time.Time `bson:"created"`
}

func (user User) ValidatePassword(password []byte) error {
	return bcrypt.CompareHashAndPassword(user.Password, password)
}
