// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["archipel_admin_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["archipel_admin_api/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["archipel_admin_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["archipel_admin_api/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["archipel_admin_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["archipel_admin_api/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["archipel_admin_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["archipel_admin_api/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["archipel_admin_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["archipel_admin_api/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["archipel_admin_api/controllers:UserController"] = append(beego.GlobalControllerRouter["archipel_admin_api/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["archipel_admin_api/controllers:UserController"] = append(beego.GlobalControllerRouter["archipel_admin_api/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["archipel_admin_api/controllers:UserController"] = append(beego.GlobalControllerRouter["archipel_admin_api/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["archipel_admin_api/controllers:UserController"] = append(beego.GlobalControllerRouter["archipel_admin_api/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["archipel_admin_api/controllers:UserController"] = append(beego.GlobalControllerRouter["archipel_admin_api/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["archipel_admin_api/controllers:UserController"] = append(beego.GlobalControllerRouter["archipel_admin_api/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["archipel_admin_api/controllers:UserController"] = append(beego.GlobalControllerRouter["archipel_admin_api/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

}