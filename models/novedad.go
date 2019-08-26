package models

import (
	"time"

	"github.com/udistrital/novedades_api/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"
)

const NovedadCollection = "novedad"

type Novedad struct {
	Id                         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Aclaracion                 string        `json:"aclaracion" bson:"aclaracion,omitempty"`
	CamposAclaracion           []string      `json:"camposaclaracion" bson:"camposaclaracion,omitempty"`
	CamposModificacion         []string      `json:"camposmodificacion" bson:"camposmodificacion,omitempty"`
	CamposModificados          []string      `json:"camposmodificados" bson:"camposmodificados,omitempty"`
	Cedente                    int           `json:"cedente" bson:"cedente,omitempty"`
	Cesionario                 int           `json:"cesionario" bson:"cesionario,omitempty"`
	Contrato                   string        `json:"contrato" bson:"contrato,omitempty"`
	FechaAdicion               time.Time     `json:"fechaadicion" bson:"fechaadicion,omitempty"`
	FechaCesion                time.Time     `json:"fechacesion" bson:"fechacesion,omitempty"`
	FechaLiquidacion           time.Time     `json:"fechaliquidacion" bson:"fechaliquidacion,omitempty"`
	FechaProrroga              time.Time     `json:"fechaprorroga" bson:"fechaprorroga,omitempty"`
	FechaRegistro              time.Time     `json:"fecharegistro"`
	FechaReinicio              time.Time     `json:"fechareinicio" bson:"fechareinicio,omitempty"`
	FechaSolicitud             time.Time     `json:"fechasolicitud"`
	FechaSuspension            time.Time     `json:"fechasuspension" bson:"fechasuspension,omitempty"`
	FechaTerminacionAnticipada time.Time     `json:"fechaterminacionanticipada" bson:"fechaterminacionanticipada,omitempty"`
	Motivo                     string        `json:"motivo"`
	NumeroActaEntrega          int           `json:"numeroactaentrega" bson:"numeroactaentrega,omitempty"`
	NumeroCDP                  string        `json:"numerocdp" bson:"numerocdp,omitempty"`
	NumeroOficioEstadoCuentas  int           `json:"numerooficioestadocuentas" bson:"numerooficioestadocuentas,omitempty"`
	NumeroSolicitud            int           `json:"numerosolicitud"`
	Observacion                string        `json:"observacion"`
	PeriodoSuspension          int           `json:"periodosuspension" bson:"periodosuspension,omitempty"`
	PlazoActual                int           `json:"plazoactual" bson:"plazoactual,omitempty"`
	Poliza                     string        `json:"poliza" bson:"poliza,omitempty"`
	TiempoProrroga             int           `json:"tiempoprorroga" bson:"tiempoprorroga,omitempty"`
	Tiponovedad                string        `json:"tiponovedad"`
	ValorAdicion               int           `json:"valoradicion" bson:"valoradicion,omitempty"`
	ValorFinalContrato         int           `json:"valorfinalcontrato" bson:"valorfinalcontrato,omitempty"`
	Vigencia                   string        `json:"vigencia"`
}

func UpdateNovedad(session *mgo.Session, j Novedad, id string) error {
	c := db.Cursor(session, NovedadCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}

func InsertNovedad(session *mgo.Session, j Novedad) {
	c := db.Cursor(session, NovedadCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllNovedads(session *mgo.Session) []Novedad {
	c := db.Cursor(session, NovedadCollection)
	defer session.Close()
	fmt.Println("Getting all Novedads")
	var novedads []Novedad
	err := c.Find(bson.M{}).All(&novedads)
	if err != nil {
		fmt.Println(err)
	}
	return novedads
}

func GetNovedadById(session *mgo.Session, id string) ([]Novedad, error) {
	c := db.Cursor(session, NovedadCollection)
	defer session.Close()
	var novedads []Novedad
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&novedads)
	return novedads, err
}

func GetNovedadByContratoVigencia(session *mgo.Session, contrato string, vigencia string) ([]Novedad, error) {
	c := db.Cursor(session, NovedadCollection)
	defer session.Close()
	var novedads []Novedad
	err := c.Find(bson.M{"contrato": contrato, "vigencia": vigencia}).All(&novedads)
	return novedads, err
}

func GetNovedadByContratoVigenciaTipo(session *mgo.Session, contrato string, vigencia string, tiponovedad string) ([]Novedad, error) {
	c := db.Cursor(session, NovedadCollection)
	defer session.Close()
	var novedads []Novedad
	err := c.Find(bson.M{"contrato": contrato, "vigencia": vigencia, "tiponovedad": tiponovedad}).All(&novedads)
	return novedads, err
}

func DeleteNovedadById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, NovedadCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok", err
}
