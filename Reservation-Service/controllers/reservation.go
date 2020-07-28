package controllers

import (
	"Reservation-Service/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Reservation
type ReservationController struct {
	beego.Controller
}

// @Title Create
// @Description create Reservation
// @Param	body		body 	models.Reservation	true		"The Reservation content"
// @Success 200 {string} models.Reservation.Id
// @Failure 403 body is empty
// @router / [post]
func (o *ReservationController) Post() {
	var ob models.Reservation
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	Reservationid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ReservationId": Reservationid}
	o.ServeJSON()
}

// @Title Get
// @Description find Reservation by Reservationid
// @Param	ReservationId		path 	string	true		"the Reservationid you want to get"
// @Success 200 {Reservation} models.Reservation
// @Failure 403 :ReservationId is empty
// @router /:ReservationId [get]
func (o *ReservationController) Get() {
	ReservationId := o.Ctx.Input.Param(":ReservationId")
	if ReservationId != "" {
		ob, err := models.GetOne(ReservationId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all Reservations
// @Success 200 {Reservation} models.Reservation
// @Failure 403 :ReservationId is empty
// @router / [get]
func (o *ReservationController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the Reservation
// @Param	ReservationId		path 	string	true		"The Reservationid you want to update"
// @Param	body		body 	models.Reservation	true		"The body"
// @Success 200 {Reservation} models.Reservation
// @Failure 403 :ReservationId is empty
// @router /:ReservationId [put]
func (o *ReservationController) Put() {
	ReservationId := o.Ctx.Input.Param(":ReservationId")
	var ob models.Reservation
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(ReservationId, ob.Score)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the Reservation
// @Param	ReservationId		path 	string	true		"The ReservationId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 ReservationId is empty
// @router /:ReservationId [delete]
func (o *ReservationController) Delete() {
	ReservationId := o.Ctx.Input.Param(":ReservationId")
	models.Delete(ReservationId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}

