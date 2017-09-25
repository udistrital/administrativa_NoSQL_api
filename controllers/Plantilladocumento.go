package controllers

import (
	"administrativa_NoSQL_api/db"
	"administrativa_NoSQL_api/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	_ "gopkg.in/mgo.v2"
)

// Operaciones Crud PlantillaDocumento
type PlantillaDocumentoController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.PlantillaDocumento
// @Failure 403 :objectId is empty
// @router / [get]
func (j *PlantillaDocumentoController) GetAll() {
	session, _ := db.GetSession()
	obs := models.GetAllPlantillaDocumentos(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Get
// @Description get PlantillaDocumento by objectId
// @Param	objectId		path 	string	true		"El id de la PlantillaDocumento a consultar"
// @Success 200 {object} models.PlantillaDocumento
// @Failure 403 :uid is empty
// @router /:objectId [get]
func (j *PlantillaDocumentoController) Get() {
	objectId := j.GetString(":objectId")
	session, _ := db.GetSession()
	if objectId != "" {
		plantilladocumento, err := models.GetPlantillaDocumentoById(session, objectId)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = plantilladocumento
		}
	}
	j.ServeJSON()
}

// @Title Borrar PlantillaDocumento
// @Description Borrar PlantillaDocumento
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *PlantillaDocumentoController) Delete() {
	session, _ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeletePlantillaDocumentoById(session, objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear PlantillaDocumento
// @Description Crear PlantillaDocumento
// @Param	body		body 	models.PlantillaDocumento	true		"Body para la creacion de PlantillaDocumento"
// @Success 200 {int} PlantillaDocumento.Id
// @Failure 403 body is empty
// @router / [post]
func (j *PlantillaDocumentoController) Post() {
	var plantilladocumento models.PlantillaDocumento
	json.Unmarshal(j.Ctx.Input.RequestBody, &plantilladocumento)
	fmt.Println(plantilladocumento)
	session, _ := db.GetSession()
	models.InsertPlantillaDocumento(session, plantilladocumento)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the PlantillaDocumento
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *PlantillaDocumentoController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var plantilladocumento models.PlantillaDocumento
	json.Unmarshal(j.Ctx.Input.RequestBody, &plantilladocumento)
	session, _ := db.GetSession()

	err := models.UpdatePlantillaDocumento(session, plantilladocumento, objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}
