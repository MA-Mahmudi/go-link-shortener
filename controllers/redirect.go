package controllers

import (
	"context"
	"github.com/MrMohebi/golang-gin-boilerplate.git/common"
	"github.com/MrMohebi/golang-gin-boilerplate.git/models"
	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"
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
		err := models.UrlsCollection.FindOne(ctx, bson.M{"short_link": shortLink}).Decode(&url)
		if err != nil {
			common.IsErr(err, false)
			c.JSON(http.StatusNotFound, gin.H{"message": "Invalid link"})
		}

		InsertLog(c, url.Id)
		c.Redirect(http.StatusFound, url.OriginalUrl)
	}
}

func InsertLog(c *gin.Context, urlId primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	now := time.Now()

	var urlLog models.UrlLog

	urlNotFind := models.UrlsLogsCollection.FindOne(ctx, bson.M{"url_id": urlId}).Decode(&urlLog)
	if urlNotFind != nil {
		common.IsErr(urlNotFind, false)
		c.JSON(http.StatusNotFound, http.StatusNotFound)
		return
	} else {
		count, err := models.UrlsLogsCollection.CountDocuments(ctx, bson.M{"url_id": urlId})
		if err != nil {
			common.IsErr(err, false)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		userAgent := useragent.Parse(c.Request.Header["User-Agent"][0])

		isBot := 0
		if userAgent.Bot == true {
			isBot = 1
		}

		urlLog = models.UrlLog{
			Id:            primitive.NewObjectID(),
			UrlId:         urlId,
			Timestamp:     now.Unix(),
			ClickCount:    count,
			IpAddress:     c.Request.RemoteAddr,
			UserAgentOS:   userAgent.OS,
			UserAgentName: userAgent.Name,
			IsBot:         isBot,
		}

		_, err = models.UrlsLogsCollection.InsertOne(ctx, urlLog)
		if err != nil {
			common.IsErr(err, false)
			c.JSON(http.StatusInternalServerError, err)
		}
	}
}
