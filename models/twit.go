package models

type Twit struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
