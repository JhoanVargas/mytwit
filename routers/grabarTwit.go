package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JhoanVargas/mytwit/bd"
	"github.com/JhoanVargas/mytwit/models"
)

func GrabarTwit(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Twit
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTwit{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTwit(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar el twit"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logró insertar el twit", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
