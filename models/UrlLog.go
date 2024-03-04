package models

import (
	"github.com/MrMohebi/golang-gin-boilerplate.git/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlLog struct {
	Id            primitive.ObjectID `json:"id" bson:"_id"`
	UrlId         primitive.ObjectID `json:"url_id" bson:"url_id"`
	Timestamp     int64              `json:"timestamp" bson:"timestamp"`
	ClickCount    int64              `json:"click_count" bson:"click_count"`
	IpAddress     string             `json:"ip_address" bson:"ip_address"`
	UserAgentOS   string             `json:"user_agent_os" bson:"user_agent_os"`
	UserAgentName string             `json:"user_agent_name" bson:"user_agent_name"`
	IsBot         int                `json:"is_bot" bson:"is_bot"`
}

var UrlsLogsCollection *mongo.Collection = configs.GetCollection(configs.GetDBClint(), "urls_logs")
