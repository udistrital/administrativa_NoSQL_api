package controllers

import (
	"administrativa_NoSQL_api/db"
	"administrativa_NoSQL_api/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	_ "gopkg.in/mgo.v2"
)

// Operaciones Crud Actainicio
type ActainicioController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Actainicio
// @Failure 403 :objectId is empty
// @router / [get]
func (j *ActainicioController) GetAll() {
	session, _ := db.GetSession()
	obs := models.GetAllActainicios(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Get
// @Description get Actainicio by objectId
// @Param	objectId		path 	string	true		"El Id de la Actainicio a consultar"
// @Success 200 {object} models.Actainicio
// @Failure 403 :uid is empty
// @router /:objectId [get]
func (j *ActainicioController) Get() {
	objectId := j.GetString(":objectId")
	session, _ := db.GetSession()
	if objectId != "" {
		actainicio, err := models.GetActainicioById(session, objectId)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = actainicio
		}
	}
	j.ServeJSON()
}

// @Title Get Contrato Vigencia
// @Description get Actainicio by vigencia
// @Param	contrato		path 	string	true		"El contrato del Actainicio a consultar"
// @Param	vigencia		path 	string	true		"La vigencia del contrato del Actainicio a consultar"
// @Success 200 {object} models.Actainicio
// @Failure 403 :uid is empty
// @router /:contrato/:vigencia [get]
func (j *ActainicioController) GetCV() {
	contrato := j.GetString(":contrato")
	vigencia := j.GetString(":vigencia")
	session, _ := db.GetSession()
	if contrato != "" && vigencia != "" {
		actainicio, err := models.GetActainicioByContratoVigencia(session, contrato, vigencia)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = actainicio
		}
	}
	j.ServeJSON()
}

// @Title Borrar Actainicio
// @Description Borrar Actainicio
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *ActainicioController) Delete() {
	session, _ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteActainicioById(session, objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Actainicio
// @Description Crear Actainicio
// @Param	body		body 	models.Actainicio	true		"Body para la creacion de Actainicio"
// @Success 200 {int} Actainicio.Id
// @Failure 403 body is empty
// @router / [post]
func (j *ActainicioController) Post() {
	var actainicio models.Actainicio
	json.Unmarshal(j.Ctx.Input.RequestBody, &actainicio)
	fmt.Println(actainicio)
	session, _ := db.GetSession()
	models.InsertActainicio(session, actainicio)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Actainicio
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *ActainicioController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var actainicio models.Actainicio
	json.Unmarshal(j.Ctx.Input.RequestBody, &actainicio)
	session, _ := db.GetSession()

	err := models.UpdateActainicio(session, actainicio, objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}
