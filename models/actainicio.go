package models

import (
	"github.com/udistrital/administrativa_NoSQL_api/db"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"
)

const ActainicioCollection = "actainicio"

type Actainicio struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Contrato    string        `json:"contrato"`
	Vigencia    string        `json:"vigencia"`
	FechaInicio time.Time     `json:"fechainicio"`
	FechaFin    time.Time     `json:"fechafin"`
}

func UpdateActainicio(session *mgo.Session, j Actainicio, id string) error {
	c := db.Cursor(session, ActainicioCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}

func InsertActainicio(session *mgo.Session, j Actainicio) {
	c := db.Cursor(session, ActainicioCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllActainicios(session *mgo.Session) []Actainicio {
	c := db.Cursor(session, ActainicioCollection)
	defer session.Close()
	fmt.Println("Getting all Actainicios")
	var actainicios []Actainicio
	err := c.Find(bson.M{}).All(&actainicios)
	if err != nil {
		fmt.Println(err)
	}
	return actainicios
}

func GetActainicioById(session *mgo.Session, id string) ([]Actainicio, error) {
	c := db.Cursor(session, ActainicioCollection)
	defer session.Close()
	var actainicios []Actainicio
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&actainicios)
	return actainicios, err
}

func GetActainicioByContratoVigencia(session *mgo.Session, contrato string, vigencia string) ([]Actainicio, error) {
	c := db.Cursor(session, ActainicioCollection)
	defer session.Close()
	var actainicios []Actainicio
	err := c.Find(bson.M{"contrato": contrato, "vigencia": vigencia}).All(&actainicios)
	return actainicios, err
}

func DeleteActainicioById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, ActainicioCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok", err
}
