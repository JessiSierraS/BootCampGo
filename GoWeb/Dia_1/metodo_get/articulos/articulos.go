package articulos

//2. Estructura Articulos
type Articulos struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Cantidad int `json:"quantity"`
	Codigo_Valor int `json:"code_value"`
	Publicado bool `json:"is_published"`
	Vencimiento string `json:"expiration"`
	Precio float64 `json:"price"`
}

//2. Estructura Catalogo de Productos
type CatalogoProductos struct {
	Productos []Articulos
}