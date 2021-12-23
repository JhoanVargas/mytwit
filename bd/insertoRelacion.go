package bd

import (
	"context"
	"time"

	"github.com/JhoanVargas/mytwit/models"
)

func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoC.Database("mytwit")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)

	if err != nil {
		return false, err
	}
	return true, nil
}
