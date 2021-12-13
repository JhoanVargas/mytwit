package bd

import (
	"context"
	"time"

	"github.com/JhoanVargas/mytwit/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro inserta registros en la bd*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoC.Database("mytwit")
	coleccion := db.Collection("usuarios")
	u.Password, _ = EncriptarPassword(u.Password)
	result, err := coleccion.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
