package bd

import (
	"context"
	"log"
	"time"

	"github.com/JhoanVargas/mytwit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoTwits(ID string, pagina int64) ([]*models.DevuelvoTwits, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoC.Database("mytwit")
	col := db.Collection("twit")

	var resultados []*models.DevuelvoTwits

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()

	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)
	opciones.SetLimit(20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTwits
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}
