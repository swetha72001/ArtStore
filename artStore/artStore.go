package artstore

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type ArtStore struct {
	ArtID       string    `json:"artId,omitempty"  bson:"artId,omitempty"`
	ArtName     string    `json:"artName,omitempty"  bson:"artName,omitempty"`
	Artist      string    `json:"artist,omitempty"  bson:"artist,omitempty"`
	CreatedDate time.Time `json:"createdDate,omitempty"  bson:"createdDate,omitempty"`
	Hotselling  bool      `json:"hotselling,omitempty"  bson:"hotselling,omitempty"`
	PhotoURL    string    `json:"photoUrl,omitempty" bson:"photoUrl,omitempty"`
	Description string    `json:"description,omitempty"  bson:"description,omitempty"`
}

type DbArts struct {
	ArtsCollection  *mongo.Collection
	TypesCollection *mongo.Collection
}
