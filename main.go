package main

import (
	"log"

	"github.com/JhoanVargas/mytwit/bd"
	"github.com/JhoanVargas/mytwit/handlers"
)

func main() {
	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
