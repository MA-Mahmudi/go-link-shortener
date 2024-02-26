package models

import (
	"github.com/MrMohebi/golang-gin-boilerplate.git/configs"
	"go.mongodb.org/mongo-driver/mongo"
)

type Url struct {
	Id          int    `json:"id" bson:"_id"`
	Title       string `json:"title"`
	Hash        string `json:"hash" bson:"hash"`
	OriginalUrl string `json:"original_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

var UrlsCollection *mongo.Collection = configs.GetCollection(configs.GetDBClint(), "urls")
