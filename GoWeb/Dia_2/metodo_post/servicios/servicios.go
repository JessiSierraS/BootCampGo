package servicios

import (
	"metodo_post/articulos"
	"encoding/json"
	"fmt"
	"os"
)

var CatalogoProductos = articulos.CatalogoProductos {}

func LeerProductJSON() {
	data, err := os.ReadFile("products.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := json.Unmarshal(data, &CatalogoProductos.Productos); err != nil {
		fmt.Println(err)
		return
	}
}