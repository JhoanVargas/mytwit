package bd

import (
	"context"
	"time"

	"github.com/JhoanVargas/mytwit/models"
	"go.mongodb.org/mongo-driver/bson"
)

func LeoTwitsSeguidores(ID string, pagina int) ([]models.DevuelvoTwitsSeguidores, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoC.Database("mytwit")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)

	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "twit",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "twit",
		}})

	condiciones = append(condiciones, bson.M{"$unwind": "$twit"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"twit.fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)

	var result []models.DevuelvoTwitsSeguidores
	err = cursor.All(ctx, &result)

	if err != nil {
		return result, true
	}

	return result, true

}
