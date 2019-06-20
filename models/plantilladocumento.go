package models

import (
	"github.com/udistrital/administrativa_NoSQL_api/db"
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const PlantillaDocumentoCollection = "plantilladocumento"

type PlantillaDocumento struct {
	Id                bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Documentoasociado string        `json:"documentoasociado"`
	Plantilla         string        `json:"plantilla"`
	FechaCreacion     time.Time     `json:"fechacreacion"`
}

func UpdatePlantillaDocumento(session *mgo.Session, j PlantillaDocumento, id string) error {
	c := db.Cursor(session, PlantillaDocumentoCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}

func InsertPlantillaDocumento(session *mgo.Session, j PlantillaDocumento) {
	c := db.Cursor(session, PlantillaDocumentoCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllPlantillaDocumentos(session *mgo.Session) []PlantillaDocumento {
	c := db.Cursor(session, PlantillaDocumentoCollection)
	defer session.Close()
	fmt.Println("Getting all plantilladocumentos")
	var plantilladocumentos []PlantillaDocumento
	err := c.Find(bson.M{}).All(&plantilladocumentos)
	if err != nil {
		fmt.Println(err)
	}
	return plantilladocumentos
}

func GetPlantillaDocumentoById(session *mgo.Session, id string) ([]PlantillaDocumento, error) {
	c := db.Cursor(session, PlantillaDocumentoCollection)
	defer session.Close()
	var plantilladocumentos []PlantillaDocumento
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&plantilladocumentos)
	return plantilladocumentos, err
}

func DeletePlantillaDocumentoById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, PlantillaDocumentoCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok", err
}
