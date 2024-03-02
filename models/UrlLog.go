package models

import (
	"github.com/MrMohebi/golang-gin-boilerplate.git/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlLog struct {
	Id         primitive.ObjectID  `json:"id" bson:"_id"`
	UrlId      primitive.ObjectID  `json:"url_id" bson:"url_id"`
	Timestamp  primitive.Timestamp `json:"timestamp" bson:"timestamp"`
	ClickCount int                 `json:"click_count" bson:"click_count"`
	Title      int                 `json:"title" bson:"title"`
}

var UrlsLogsCollection *mongo.Collection = configs.GetCollection(configs.GetDBClint(), "urls_logs")
