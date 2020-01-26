package db

import (
	"../models"
	"../settings"
	"errors"
	"github.com/number571/gopeer"
)

func SetUser(user *models.User) error {
	if InUsers(user.Auth.Hashpasw) {
		return errors.New("User already exist")
	}
	_, err := settings.DB.Exec(
		"INSERT INTO User (Hashpasw, Key) VALUES ($1, $2)",
		user.Auth.Hashpasw,
		gopeer.Base64Encode(
			gopeer.EncryptAES(
				user.Auth.Pasw,
				[]byte(gopeer.StringPrivate(user.Keys.Private)),
			),
		),
	)
	if err != nil {
		panic("exec 'setuser' failed")
	}
	return nil
}