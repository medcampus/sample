package db

import (
	"github.com/medcampus/backend/clap/utils"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Handler struct {
	C *mgo.Collection
}

func NewHandler(session *mgo.Session) *Handler {
	return &Handler{
		session.DB(viper.GetString("mongo.dbName")).C(viper.GetString("mongo.collectionName")),
	}
}

func (h *Handler) AddClap(serviceId string) (claps int64, err error) {
	var clap Clap

	clapId := bson.ObjectIdHex(serviceId)

	err = h.C.Find(bson.M{"clapid": clapId}).One(&clap)
	if err != nil {
		if err.Error() == utils.ErrNotFound.Error() {
			h.C.Insert(&Clap{
				clapId,
				1,
			})

			return 1, nil
		}

		return 0, err
	}

	clap.Claps +=1

	err = h.C.Update(bson.M{"clapid": clapId}, bson.M{"$set": bson.M{"claps": clap.Claps}})

	if err != nil {
		return 0, err
	}

	return clap.Claps, nil
}

func (h *Handler) GetClaps(serviceId string) (claps int64, err error){

	var clap Clap

	err = h.C.Find(bson.M{"clapid": bson.ObjectIdHex(serviceId)}).One(&clap)
	if err != nil {
		return 0, err
	}

	return clap.Claps, nil
}