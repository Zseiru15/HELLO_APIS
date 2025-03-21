// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/sena_2824182/mid_prueba/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/buscarcarta", //localhost:8081/v1/buscarcarta/"#"
			beego.NSInclude(
				&controllers.Carta_contadorController{},
			),
		),
		beego.NSNamespace("/buscarpaises", //http://localhost:8081/v1/buscarpaises/capitales/"Pais"
			beego.NSInclude(
				&controllers.PaisesController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
