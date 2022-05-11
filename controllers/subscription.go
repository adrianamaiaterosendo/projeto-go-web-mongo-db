package controllers

import (
	"github.com/adrianamaiaterosendo/projeto-go-web-mongo-db.git/db"
	"github.com/adrianamaiaterosendo/projeto-go-web-mongo-db.git/model"
)

func Create(name, email string) error {
	s := model.Subscription{Name: name, Email: email}

	return db.Insert("newsletter", s)
}
