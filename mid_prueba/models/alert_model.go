package models

type Alert struct {
	Type string
	Code string
	Body interface{}
}

// Estructura para los países (Paises)
type Paises struct {
	Nombre    string
	Poblacion int
	Fronteras []interface{}
	Fifa      string
}
