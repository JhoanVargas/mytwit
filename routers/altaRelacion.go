package routers

import (
	"net/http"

	"github.com/JhoanVargas/mytwit/bd"
	"github.com/JhoanVargas/mytwit/models"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al insertar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se logro insertar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
