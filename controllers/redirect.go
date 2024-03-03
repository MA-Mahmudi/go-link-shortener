package controllers

import (
	"context"
	"fmt"
	"github.com/MrMohebi/golang-gin-boilerplate.git/common"
	"github.com/MrMohebi/golang-gin-boilerplate.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func RedirectUrl() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		shortLink := c.Param("short_url")
		var url models.Url
		err := models.UrlsCollection.FindOne(ctx, bson.D{{"short_link", shortLink}}).Decode(&url)
		common.IsErr(err, true)

		InsertLog(c, url.Id)
		c.Redirect(http.StatusFound, url.OriginalUrl)
	}
}

func InsertLog(c *gin.Context, urlId primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	now := time.Now()

	count, err := models.UrlsLogsCollection.CountDocuments(ctx, bson.M{"url_id": urlId})
	fmt.Println(count)

	var urlLog models.UrlLog

	urlNotFind := models.UrlsLogsCollection.FindOne(ctx, bson.M{"url_id": urlId}).Decode(&urlLog)
	fmt.Println(urlLog)
	if urlNotFind != nil {
		common.IsErr(urlNotFind, false)
		c.JSON(http.StatusNotFound, gin.H{"message": "لینک مورد نظر یافت نشد."})
	} else {
		urlLog.Id = primitive.NewObjectID()
		urlLog.ClickCount = count
		urlLog.Timestamp = now.Unix()
		_, err = models.UrlsLogsCollection.InsertOne(ctx, urlLog)
		if err != nil {
			common.IsErr(err, false)
			c.JSON(http.StatusInternalServerError, err)
		}
	}
}
