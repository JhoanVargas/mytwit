package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/JhoanVargas/mytwit/bd"
	"github.com/JhoanVargas/mytwit/models"
)

func SubirAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

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

	usuario.Avatar = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)

	if err != nil || status == false {
		http.Error(w, "Error al subir el avatar a la Base de Datos"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "apliccation/json")
	w.WriteHeader(http.StatusCreated)
}
