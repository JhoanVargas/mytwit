package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BorroTwit(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoC.Database("mytwit")
	col := db.Collection("twit")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(ctx, condicion)
	return err

}
