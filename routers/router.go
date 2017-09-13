// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"administrativa_NoSQL_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/actainicio",
			beego.NSInclude(
				&controllers.ActainicioController{},
			),
		),

		beego.NSNamespace("/tiponovedad",
			beego.NSInclude(
				&controllers.TiponovedadController{},
			),
		),

		beego.NSNamespace("/novedad",
			beego.NSInclude(
				&controllers.NovedadController{},
			),
		),

		beego.NSNamespace("/plantilladocumento",
			beego.NSInclude(
				&controllers.PlantillaDocumentoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
