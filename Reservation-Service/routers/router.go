package routers

import (
	"Reservation-Service/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSCond(func(ctx *context.Context) bool {
    	    if ctx.Input.Domain() == "localhost" {
    	        return true
    	    }
    	    return false
    	}),
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
	beego.Router("/", &controllers.ReservationController{})
	

}
