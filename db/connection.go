package db

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

var (
	session  *mgo.Session
	mongoUrl string
)

// create root DB Session
func CreateDBSession() {
	var err error
	mongoUrl = viper.GetString("mongo.host")
	session, err = mgo.Dial(mongoUrl)
	if err != nil {
		log.Errorf("Unable to create ROOT database session, Error %v", err)
		return
	}

	log.Infof("Created ROOT Database Session from %s", mongoUrl)
}

// copy database session
func GetMongoSession() *mgo.Session {
	if session == nil {
		CreateDBSession()
	}

	log.Info("Copied mongo session from root session")
	return session.Copy()
}
