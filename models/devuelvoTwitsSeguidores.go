package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevuelvoTwitsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitemprty"`
	UsuarioID         string             `bson:"ususarioid" json:"userId,omitemprty"`
	UsuarioRelacionID string             `bson:"ususariorelacionid" json:"userRelationId,omitemprty"`
	Twit              struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitemprty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitemprty"`
		ID      string    `bson:"_id" json:"_id,omitemprty"`
	}
}
