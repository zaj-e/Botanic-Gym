package routers

import (
	"Reservation-Service/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/reservations",
			beego.NSInclude(
				&controllers.ReservationController{},
			),
		),		
	)
	beego.AddNamespace(ns)
    // beego.Router("/", &controllers.MainController{})
}
