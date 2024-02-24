package models

import (
	"github.com/MrMohebi/golang-gin-boilerplate.git/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlLog struct {
	UrlId      primitive.ObjectID `json:"url_id"`
	ClickCount int                `json:"click_count"`
	UserAgent  string             `json:"user_agent"`
}

var UrlsLogsCollection *mongo.Collection = configs.GetCollection(configs.GetDBClint(), "urls_logs")
