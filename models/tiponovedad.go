package models

import (
	"administrativa_NoSQL_api/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"
)

const TiponovedadCollection = "tiponovedad"

type Tiponovedad struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Nombre      string        `json:"nombre"`
	Descripcion string        `json:"descripcion"`
}

func UpdateTiponovedad(session *mgo.Session, j Tiponovedad, id string) error {
	c := db.Cursor(session, TiponovedadCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}

func InsertTiponovedad(session *mgo.Session, j Tiponovedad) {
	c := db.Cursor(session, TiponovedadCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllTiponovedads(session *mgo.Session) []Tiponovedad {
	c := db.Cursor(session, TiponovedadCollection)
	defer session.Close()
	fmt.Println("Getting all Tiponovedads")
	var tiponovedads []Tiponovedad
	err := c.Find(bson.M{}).All(&tiponovedads)
	if err != nil {
		fmt.Println(err)
	}
	return tiponovedads
}

func GetTiponovedadById(session *mgo.Session, id string) ([]Tiponovedad, error) {
	c := db.Cursor(session, TiponovedadCollection)
	defer session.Close()
	var tiponovedads []Tiponovedad
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&tiponovedads)
	return tiponovedads, err
}

func DeleteTiponovedadById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, TiponovedadCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok", err
}
