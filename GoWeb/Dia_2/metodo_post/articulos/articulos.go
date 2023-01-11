package articulos

//2. Estructura Articulos //binding para validar vacio
type Articulos struct {
	ID 				int 	`json:"id"`
	Nombre 			string	`json:"name" binding:"required"`
	Cantidad 		int 	`json:"quantity" binding:"required"`
	Codigo_Valor 	string 	`json:"code_value" binding:"required"`
	Publicado 		bool 	`json:"is_published"`
	Vencimiento 	string 	`json:"expiration" binding:"required"`
	Precio 			float64 `json:"price" binding:"required"`
}

//2. Estructura Catalogo de Productos
type CatalogoProductos struct {
	Productos []Articulos
}