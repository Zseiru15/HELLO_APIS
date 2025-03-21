package controllers

import (
	//"fmt"

	"encoding/json"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/mid_prueba/models"
	"github.com/sena_2824182/mid_prueba/services"
)

// Carta_contadorController operations for Carta_contador
type Carta_contadorController struct {
	beego.Controller
}

// URLMapping ...
func (c *Carta_contadorController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Carta_contador
// @Param	body		body 	models.Carta_contador	true		"body for Carta_contador content"
// @Success 201 {object} models.Carta_contador
// @Failure 403 body is empty
// @router / [post]
func (c *Carta_contadorController) Post() {

	var body_ingresa []map[string]interface{}
	var alerta models.Alert
	var temporal_usuario []byte
	var temporal_producto []byte

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa); err == nil {
		//fmt.Println("body que ingresa:", body_ingresa)

		jsonData, err := json.MarshalIndent(body_ingresa, "", " ")
		if err != nil {
			fmt.Println("Error al convertir a JSON", err)
		}

		json_usuario := body_ingresa[0]
		json_producto := body_ingresa[1]
		//fmt.Println("Body usuario:", json_usuario)
		//fmt.Println("Body producto:", json_producto)

		json_usuario_byte, _ := json.Marshal(json_usuario)
		response_usuario, _ := services.Metodo_post("Servicio_post", json_usuario_byte)

		json_producto_byte, _ := json.Marshal(json_producto)
		response_producto, _ := services.Metodo_post("Servicio_post", json_producto_byte)

		temporal_usuario = response_usuario
		temporal_producto = response_producto

		//fmt.Println("Esto responde el post de usuario:", string(response_usuario))
		//fmt.Println("Esto respnde el post de producto:", string(response_producto))

		fmt.Print("Body de ingreso en JSON:", string(jsonData))
	}
	var temporal_usuario2 map[string]interface{}
	var temporal_producto2 map[string]interface{}

	json.Unmarshal(temporal_usuario, &temporal_usuario2)
	json.Unmarshal(temporal_producto, &temporal_producto2)
	var body_final []map[string]interface{}
	body_final = append(body_final, temporal_usuario2["data"].(map[string]interface{}))
	body_final = append(body_final, temporal_producto2["data"].(map[string]interface{}))

	alerta.Code = "201"
	alerta.Type = "post"
	alerta.Body = body_final
	c.Data["json"] = alerta
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Carta_contador by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param id_2 path int true "este es un segundo parametro"
// @Success 200 {object} models.Carta_contador
// @Failure 403 :id is empty
// @router /:id/:id_2 [get]
func (c *Carta_contadorController) GetOne() {
	//fmt.Println("Funcion Get")
	id_ingreso := c.Ctx.Input.Param(":id") // para capturar el parametro del url /id
	//id_ingreso_2 := c.Ctx.Input.Param(":id_2")
	//fmt.Println("EL primer id de ingreso es:", id_ingreso)
	//fmt.Println("El segundo Id de ingreso es:", id_ingreso_2)

	//asignacion de datos al body
	body, _ := services.Metodo_get("Servicio_Cartas", id_ingreso)
	body2, _ := services.Metodo_get("Servicio_Usuarios", id_ingreso)

	//resp, err := http.Get(beego.AppConfig.String("Servicio_Cartas")+ id_ingreso)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer resp.Body.Close()

	//resp2, err2 := http.Get(beego.AppConfig.String("Servicio_Usuarios")+ id_ingreso)
	//if err2 != nil {
	//	log.Fatal(err2)
	//}
	//defer resp2.Body.Close()
	//----------------------------------------------------------------------------------------

	// Leer la respuesta
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Leer la respuesta
	//body2, err2 := ioutil.ReadAll(resp2.Body)
	//if err2 != nil {
	//	log.Fatal(err2)
	//}
	//----------------------------------------------------------------------------------------

	// Imprimir el resultado
	//fmt.Println("Status Code:", resp.Status)
	//fmt.Println("Response Body:", string(body))
	//fmt.Println("Response Body:", string(body2))

	resultado1, _ := services.ProcesarJsonArreglos(body)
	//----------------------------------------------------------------------------------------

	//var result map[string]interface{}  // El JSON que esperas es un array de objetos
	var result2 map[string]interface{} // El JSON que esperas es un array de objetos

	//err = json.Unmarshal(body, &result)
	//if err != nil {
	//	log.Fatal("Error al parsear JSON:", err)
	//}

	err2 := json.Unmarshal(body2, &result2)
	if err2 != nil {
		log.Fatal("Error al parsear JSON:", err2)
	}

	//agregar un campo nuevo
	for i := range resultado1 {
		resultado1[i] = map[string]interface{}{
			"Campo_nuevo": i + 1,
			"body":        resultado1[i]["body"],
		}
	}

	result2 = map[string]interface{}{
		"direccion":          result2["address"],
		"telefono":           result2["phone"],
		"codigo_postal":      result2["address"].(map[string]interface{})["zipcode"],
		"dirrecion_telefono": map[string]interface{}{"dirrecion": result2["address"], "telefono": result2["phone"]},
	}

	//------------------------------------------------

	//Sacar una parte de un json del resultado 1
	resultado := append(resultado1, result2)

	//informacion de estado
	//fmt.Println("La cantidad de datos son", len(resultado))
	c.Data["json"] = map[string]interface{}{
		"Succes":          true,
		"Status":          200,
		"Message":         "Consulta existosa",
		"Data":            resultado,
		"Cantidad Cartas": len(resultado)}
	c.ServeJSON()

}

// GetAll ...
// @Title GetAll
// @Description get Carta_contador
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Carta_contador
// @Failure 403
// @router / [get]
func (c *Carta_contadorController) GetAll() {

}

//----------------------------------------------------------------------------------------------------------------------------------------

// Put ...
// @Title Put
// @Description update the Carta_contador
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Carta_contador	true		"body for Carta_contador content"
// @Success 200 {object} models.Carta_contador
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Carta_contadorController) Put() {
	var body_ingresa map[string]interface{}
	var response map[string]interface{}

	id_ingreso := c.Ctx.Input.Param(":id") // para capturar el parametro del url /id

	//fmt.Println("Ojo entra al put")

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa); err == nil {
		fmt.Println("Este es el id ingresado:", id_ingreso)
		fmt.Println("Este es el body ingresado:", body_ingresa)
		get_inicial, _ := services.Metodo_get("Servicio_Cartas2", id_ingreso)

		//fmt.Println("Este es el get:", string(get_inicial))

		json_inicial, _ := services.ProcesarJson(get_inicial)

		//fmt.Println("Este es el json:", json_inicial)

		nuevo_json := map[string]interface{}{
			"userId": json_inicial["userId"],
			"id":     json_inicial["id"],
			"title":  body_ingresa["title"],
			"body":   body_ingresa["body"],
		}
		//fmt.Println("Este es el nuevo json:", nuevo_json)

		json_byte, _ := json.Marshal(nuevo_json)

		response_put, _ := services.Metodo_put("Servicio_Cartas2", id_ingreso, json_byte)

		//fmt.Println("Este es el response del put:", string(response_put))
		response, _ = services.ProcesarJson(response_put)
	}

	c.Data["json"] = map[string]interface{}{
		"Succes":  true,
		"Status":  200,
		"Message": "Actualizacion existosa",
		"Data":    response,
	}
	c.ServeJSON()

}

// Delete ...
// @Title Delete
// @Description delete the Carta_contador
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Carta_contadorController) Delete() {

	id_ingreso := c.Ctx.Input.Param(":id") // para capturar el parametro del url /id
	fmt.Println(id_ingreso)
	get_inicial, _ := services.Metodo_get("Servicio_Cartas2", id_ingreso)
	fmt.Println("Este es el get inicial", string(get_inicial))
	json_inicial, _ := services.ProcesarJson(get_inicial)

	json_nuevo := map[string]interface{}{
		"userId": json_inicial["userId"],
		"id":     json_inicial["id"],
		"title":  json_inicial["title"],
		"body":   "false",
	}
	fmt.Println("json nuevo:", json_nuevo)

	json_byte, _ := json.Marshal(json_nuevo)

	response_put, _ := services.Metodo_put("Servicio_Cartas2", id_ingreso, json_byte)

	fmt.Println("Este es el response del put:", string(response_put))

	jsonpost := map[string]interface{}{
		"userId": json_inicial["userId"],
		"id":     json_inicial["id"],
		"title":  json_inicial["title"],
		"body":   "false",
	}

	json_byte_post, _ := json.Marshal(jsonpost)

	response_post, _ := services.Metodo_post("Servicio_Cartas2", json_byte_post)

	fmt.Println("Este es el post:", string(response_post))

	c.Data["json"] = map[string]interface{}{
		"succes":    "true",
		"code":      200,
		"status ok": "ok",
		"data":      "Registro eliminado",
	}
	c.ServeJSON()

}
