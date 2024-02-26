package controllers

import (
	"context"
	"github.com/MrMohebi/golang-gin-boilerplate.git/common"
	"github.com/MrMohebi/golang-gin-boilerplate.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

var urlsColl = models.UrlsCollection
var now = time.Now()

func CreateUrl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody models.Url
		err := ctx.BindJSON(&reqBody)
		common.IsErr(err, true)

		shortUrl := common.RandStr(5)

		url := models.Url{
			Id:          primitive.NewObjectID(),
			Title:       reqBody.Title,
			ShortLink:   shortUrl,
			OriginalUrl: reqBody.OriginalUrl,
			CreatedAt:   now.Format(time.DateTime),
			UpdatedAt:   now.Format(time.DateTime),
		}

		_, err = urlsColl.InsertOne(context.TODO(), url)
		common.IsErr(err, true)

		ctx.JSON(http.StatusOK, url)
	}
}

func RedirectUrl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		shortLink := ctx.Param("short_url")
		var url models.Url
		err := urlsColl.FindOne(context.TODO(), bson.D{{"short_link", shortLink}}).Decode(&url)
		common.IsErr(err, true)
		ctx.Redirect(302, url.OriginalUrl)
	}
}
