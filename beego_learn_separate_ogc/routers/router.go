// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Go_learn/beego_learn_separate_ogc/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/t_lesson",
			beego.NSInclude(
				&controllers.TLessonController{},
			),
		),

		beego.NSNamespace("/t_match",
			beego.NSInclude(
				&controllers.TMatchController{},
			),
		),

		beego.NSNamespace("/t_options",
			beego.NSInclude(
				&controllers.TOptionsController{},
			),
		),

		beego.NSNamespace("/t_organization",
			beego.NSInclude(
				&controllers.TOrganizationController{},
			),
		),

		beego.NSNamespace("/t_project",
			beego.NSInclude(
				&controllers.TProjectController{},
			),
		),

		beego.NSNamespace("/t_stu_lesson",
			beego.NSInclude(
				&controllers.TStuLessonController{},
			),
		),

		beego.NSNamespace("/t_user_match",
			beego.NSInclude(
				&controllers.TUserMatchController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
