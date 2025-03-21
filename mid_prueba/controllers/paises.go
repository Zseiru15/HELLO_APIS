package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/mid_prueba/models"
	"github.com/sena_2824182/mid_prueba/services"
)

// PaisesController operations for Paises
type PaisesController struct {
	beego.Controller
}

// URLMapping ...
func (c *PaisesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Paises
// @Param	body		body 	models.Paises	true		"body for Paises content"
// @Success 201 {object} models.Paises
// @Failure 403 body is empty
// @router / [post]
func (c *PaisesController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Paises by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Paises
// @Failure 403 :id is empty
// @router /capitales/:capital1/:capital2/:capital3/:capital4/:capital5 [get]
func (c *PaisesController) GetOne() {
	fmt.Println("Funcion Get")

	// Obtener los parámetros del URL de forma dinámica
	capitales := []string{
		c.Ctx.Input.Param(":capital1"),
		c.Ctx.Input.Param(":capital2"),
		c.Ctx.Input.Param(":capital3"),
		c.Ctx.Input.Param(":capital4"),
		c.Ctx.Input.Param(":capital5"),
	}

	// Imprimir los parámetros de las capitales
	for i, capital := range capitales {
		fmt.Printf("Capital %d: %s\n", i+1, capital)
	}

	// Obtener datos de los países por cada capital
	var resultados []map[string]interface{}
	for _, capital := range capitales {
		body, err := services.Metodo_get("Servicio_Capitales", capital)
		if err != nil {
			fmt.Printf("Error al obtener datos de la capital %s: %v\n", capital, err)
			continue
		}

		fmt.Println("Response Body:", string(body))

		// Procesar la respuesta JSON
		resultado, err := services.ProcesarJsonArreglos(body)
		if err != nil {
			fmt.Printf("Error al procesar la respuesta JSON para la capital %s: %v\n", capital, err)
			continue
		}

		// Agregar los resultados procesados
		resultados = append(resultados, resultado...)
	}

	// Tasa de cambio fija para COP a otras divisas
	EUR := 0.00022 // 1 COP = 0.00022 EUR
	USD := 0.00026 // 1 COP = 0.00026 USD
	RUB := 0.024   // 1 COP = 0.024 RUB
	MXN := 0.048   // 1 COP = 0.048 MXN
	JPY := 0.028   // 1 COP = 0.028 JPY

	// Valor fijo de COP (1 COP)
	copValue := 1.0

	// Realizar la conversión
	for i := range resultados {
		// Accedemos al campo "currencies" que es un mapa
		moneda := resultados[i]["currencies"].(map[string]interface{})

		// Iterar sobre las claves del mapa "moneda"
		for key := range moneda {
			currencyKey := key
			// fmt.Println("Clave encontrada:", currencyKey)

			// Crear el mapa de cambio vacío
			cambio := map[string]float64{"COP": copValue}

			// Realizamos la conversión dependiendo de la moneda
			if currencyKey == "COP" {
				cambio["COP"] = copValue
			} else if currencyKey == "EUR" {
				cambio["EUR"] = copValue * EUR
			} else if currencyKey == "USD" {
				cambio["USD"] = copValue * USD
			} else if currencyKey == "RUB" {
				cambio["RUB"] = copValue * RUB
			} else if currencyKey == "MXN" {
				cambio["MXN"] = copValue * MXN
			} else if currencyKey == "JPY" {
				cambio["JPY"] = copValue * JPY
			}

			// Ahora añadimos el resultado al mapa final
			resultados[i] = map[string]interface{}{
				"pais":           resultados[i]["name"].(map[string]interface{})["common"],
				"nombre_oficial": resultados[i]["name"].(map[string]interface{})["official"],
				"capital":        resultados[i]["capital"],
				"moneda":         resultados[i]["currencies"],
				"cambio":         cambio,
				"region":         resultados[i]["region"],
				"subregion":      resultados[i]["subregion"],
			}
		}
	}

	// Devolver la respuesta en formato JSON
	c.Data["json"] = map[string]interface{}{
		"Succes":   true,
		"Status":   200,
		"Message":  "Consulta exitosa",
		"Data":     resultados,
		"Cantidad": len(resultados),
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Paises
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Paises
// @Failure 403
// @router / [get]
func (c *PaisesController) GetAll() {
	// Obtener el parámetro de subregión de la URL
	subregion := c.Ctx.Input.Param(":subregion")
	log.Println("Obteniendo países de la subregión:", subregion)

	// Codificar el parámetro de la subregión para la URL (reemplazar los espacios por '%20')
	url := fmt.Sprintf("https://restcountries.com/v3.1/subregion/%s", subregion)
	resp, err := http.Get(url)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Error al hacer la solicitud a la API externa"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Error al leer la respuesta de la API externa"}
		c.ServeJSON()
		return
	}

	// Imprimir la respuesta para depuración
	log.Println("Cuerpo de la respuesta de la API:", string(body))

	// Convertir la respuesta JSON en un objeto
	var paises []map[string]interface{}
	err = json.Unmarshal(body, &paises)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Error al parsear la respuesta JSON"}
		c.ServeJSON()
		return
	}

	// Crear una lista para almacenar los países
	var resultado []models.Paises

	// Extraer la información de cada país
	for _, p := range paises {
		nombre, _ := p["name"].(map[string]interface{})["common"].(string)
		poblacion, _ := p["population"].(float64)
		fronteras, _ := p["borders"].([]interface{})
		fifaCode := ""
		if code, ok := p["fifa"].(string); ok {
			fifaCode = code
		}

		// Crear el objeto para el país
		resultado = append(resultado, models.Paises{
			Nombre:    nombre,
			Poblacion: int(poblacion),
			Fronteras: fronteras,
			Fifa:      fifaCode,
		})
	}

	// Responder con los datos de todos los países
	c.Data["json"] = resultado
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Paises
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Paises	true		"body for Paises content"
// @Success 200 {object} models.Paises
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PaisesController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Paises
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PaisesController) Delete() {

}
