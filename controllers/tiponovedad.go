package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/udistrital/novedades_api/db"
	"github.com/udistrital/novedades_api/models"

	"github.com/astaxie/beego"
	_ "gopkg.in/mgo.v2"
)

// Operaciones Crud Tiponovedad
type TiponovedadController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Tiponovedad
// @Failure 403 :objectId is empty
// @router / [get]
func (j *TiponovedadController) GetAll() {
	session, _ := db.GetSession()
	obs := models.GetAllTiponovedads(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Get
// @Description get Tiponovedad by id
// @Param	objectId		path 	string	true		"El id de la Tiponovedad a consultar"
// @Success 200 {object} models.Tiponovedad
// @Failure 403 :uid is empty
// @router /:objectId [get]
func (j *TiponovedadController) Get() {
	objectId := j.Ctx.Input.Param(":objectId")
	session, _ := db.GetSession()
	if objectId != "" {
		tiponovedad, err := models.GetTiponovedadById(session, objectId)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = tiponovedad
		}
	}
	j.ServeJSON()
}

// @Title Borrar Tiponovedad
// @Description Borrar Tiponovedad
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *TiponovedadController) Delete() {
	session, _ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteTiponovedadById(session, objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Post Tiponovedad
// @Description Crear Tiponovedad
// @Param	body		body 	models.Tiponovedad	true		"Body para la creacion de Tiponovedad"
// @Success 200 {int} Tiponovedad.Id
// @Failure 403 body is empty
// @router / [post]
func (j *TiponovedadController) Post() {
	var tiponovedad models.Tiponovedad
	json.Unmarshal(j.Ctx.Input.RequestBody, &tiponovedad)
	fmt.Println(tiponovedad)
	session, _ := db.GetSession()
	models.InsertTiponovedad(session, tiponovedad)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Tiponovedad
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *TiponovedadController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var tiponovedad models.Tiponovedad
	json.Unmarshal(j.Ctx.Input.RequestBody, &tiponovedad)
	session, _ := db.GetSession()

	err := models.UpdateTiponovedad(session, tiponovedad, objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}
