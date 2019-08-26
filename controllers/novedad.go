package controllers

import (
	"github.com/udistrital/novedades_api/models"
	"github.com/udistrital/novedades_api/db"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

// Operaciones Crud Novedad
type NovedadController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Novedad
// @Failure 403 :objectId is empty
// @router / [get]
func (j *NovedadController) GetAll() {
	session, _ := db.GetSession()
	obs := models.GetAllNovedads(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Get
// @Description get Novedad by objectId
// @Param	objectId		path 	string	true		"El objectId de la Novedad a consultar"
// @Success 200 {object} models.Novedad
// @Failure 403 :uid is empty
// @router /:objectId [get]
func (j *NovedadController) Get() {
	objectId := j.GetString(":objectId")
	session, _ := db.GetSession()
	if objectId != "" {
		novedad, err := models.GetNovedadById(session, objectId)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = novedad
		}
	}
	j.ServeJSON()
}

// @Title Get Contrato Vigencia
// @Description get Novedad by vigencia
// @Param	contrato		path 	string	true		"El contrato de la Novedad a consultar"
// @Param	vigencia		path 	string	true		"La vigencia del contrato de la Novedad a consultar"
// @Success 200 {object} models.Novedad
// @Failure 403 :uid is empty
// @router /:contrato/:vigencia [get]
func (j *NovedadController) GetCV() {
	contrato := j.GetString(":contrato")
	vigencia := j.GetString(":vigencia")
	session, _ := db.GetSession()
	if contrato != "" && vigencia != "" {
		novedad, err := models.GetNovedadByContratoVigencia(session, contrato, vigencia)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = novedad
		}
	}
	j.ServeJSON()
}

// @Title Get Contrato Vigencia tiponovedad
// @Description get Novedad by contrato vigencia tiponovedad
// @Param	contrato		path 	string	true		"El contrato de la Novedad a consultar"
// @Param	vigencia		path 	string	true		"La vigencia del contrato de la Novedad a consultar"
// @Param	tiponovedad		path 	string	true		"El tipo de la Novedad a consultar"
// @Success 200 {object} models.Novedad
// @Failure 403 :uid is empty
// @router /:contrato/:vigencia/:tiponovedad [get]
func (j *NovedadController) GetCVT() {
	contrato := j.GetString(":contrato")
	vigencia := j.GetString(":vigencia")
	tiponovedad := j.GetString(":tiponovedad")
	session, _ := db.GetSession()
	if contrato != "" && vigencia != "" {
		novedad, err := models.GetNovedadByContratoVigenciaTipo(session, contrato, vigencia, tiponovedad)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = novedad
		}
	}
	j.ServeJSON()
}

// @Title Borrar Novedad
// @Description Borrar Novedad
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *NovedadController) Delete() {
	session, _ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteNovedadById(session, objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Novedad
// @Description Crear Novedad
// @Param	body		body 	models.Novedad	true		"Body para la creacion de Novedad"
// @Success 200 {int} Novedad.Id
// @Failure 403 body is empty
// @router / [post]
func (j *NovedadController) Post() {
	var novedad models.Novedad
	json.Unmarshal(j.Ctx.Input.RequestBody, &novedad)
	fmt.Println(novedad)
	session, _ := db.GetSession()
	models.InsertNovedad(session, novedad)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Novedad
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *NovedadController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var novedad models.Novedad
	json.Unmarshal(j.Ctx.Input.RequestBody, &novedad)
	session, _ := db.GetSession()

	err := models.UpdateNovedad(session, novedad, objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}
