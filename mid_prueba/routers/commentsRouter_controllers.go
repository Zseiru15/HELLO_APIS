package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:Carta_contadorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:Carta_contadorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:Carta_contadorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:Carta_contadorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:Carta_contadorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:Carta_contadorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:Carta_contadorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:Carta_contadorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:Carta_contadorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:Carta_contadorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id/:id_2",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:PaisesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:PaisesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:PaisesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:PaisesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:PaisesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:PaisesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:PaisesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:PaisesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:PaisesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/mid_prueba/controllers:PaisesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/capitales/:capital1/:capital2/:capital3/:capital4/:capital5",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
