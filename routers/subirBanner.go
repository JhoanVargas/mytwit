package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/JhoanVargas/mytwit/bd"
	"github.com/JhoanVargas/mytwit/models"
)

func SubirBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_RDONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir la imagen!"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error al copiar la imagen!"+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)

	if err != nil || status == false {
		http.Error(w, "Error al subir el banner a la Base de Datos"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "apliccation/json")
	w.WriteHeader(http.StatusCreated)
}
