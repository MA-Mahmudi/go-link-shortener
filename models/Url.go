package models

import (
	"github.com/MrMohebi/golang-gin-boilerplate.git/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Url struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title"`
	Hash        string             `json:"hash" bson:"hash"`
	OriginalUrl string             `json:"original_url"`
	CreatedAt   primitive.DateTime `json:"created_at"`
	UpdatedAt   primitive.DateTime `json:"updated_at"`
	DeletedAt   primitive.DateTime `json:"deleted_at"`
}

var UrlsCollection *mongo.Collection = configs.GetCollection(configs.GetDBClint(), "urls")
