package models

import (
	"github.com/MrMohebi/golang-gin-boilerplate.git/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Url struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	ShortLink   string             `json:"short_link" bson:"short_link"`
	OriginalUrl string             `json:"original_url" bson:"original_url" validate:"required"`
	CreatedAt   string             `json:"created_at" bson:"created_at"`
	UpdatedAt   string             `json:"updated_at" bson:"updated_at"`
	DeletedAt   string             `json:"deleted_at" bson:"deleted_at"`
	ExpireDate  string             `json:"expire_date" bson:"expire_date"`
	Status      int                `json:"status" bson:"status"`
}

var UrlsCollection *mongo.Collection = configs.GetCollection(configs.GetDBClint(), "urls")
