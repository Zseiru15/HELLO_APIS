package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:AdminController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:AdminController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:ClienteController"] = append(beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:ClienteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:ClienteController"] = append(beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:ClienteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:ClienteController"] = append(beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:ClienteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:ClienteController"] = append(beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:ClienteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:ClienteController"] = append(beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:ClienteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:RolController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:RolController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:RolController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:RolController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/Zseiru15/API_CRUD_GO/controllers:RolController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
