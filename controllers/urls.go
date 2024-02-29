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

var urlsColl = models.UrlsCollection
var urlsLogColl = models.UrlsLogsCollection
var now = time.Now()

func CreateUrl() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var reqBody models.Url

		err := c.BindJSON(&reqBody)
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

		urlLog := models.UrlLog{
			Id:         primitive.NewObjectID(),
			UrlId:      url.Id,
			ClickCount: 0,
		}

		_, err = urlsColl.InsertOne(ctx, url)
		common.IsErr(err, true)

		_, err = urlsLogColl.InsertOne(ctx, urlLog)
		common.IsErr(err, true)

		c.JSON(http.StatusOK, url)
	}
}

func RedirectUrl() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		shortLink := c.Param("short_url")
		var url models.Url
		err := urlsColl.FindOne(ctx, bson.D{{"short_link", shortLink}}).Decode(&url)
		common.IsErr(err, true)
		UpdateLogs(url.Id, c)
		c.Redirect(http.StatusFound, url.OriginalUrl)

	}
}

func UpdateLogs(urlId primitive.ObjectID, c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var urlLog models.UrlLog

	urlNotFind := urlsLogColl.FindOne(ctx, bson.M{"url_id": urlId}).Decode(&urlLog)
	if urlNotFind != nil {
		common.IsErr(urlNotFind, false)
		c.JSON(http.StatusNotFound, gin.H{"message": "لینک مورد نظر یافت نشد."})
	}
	fmt.Println(urlLog)

	res, err := urlsLogColl.UpdateOne(ctx, bson.D{{"_id", urlLog.Id}}, bson.M{"$set": bson.M{"click_count": urlLog.ClickCount + 1}})
	common.IsErr(err, true)

	fmt.Println(res)
}

func GetAllUrls() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var urls []models.Url

		count, _ := urlsColl.CountDocuments(ctx, bson.M{})
		res, err := urlsColl.Find(ctx, bson.M{})
		if err != nil {
			common.IsErr(err, false)
			c.JSON(http.StatusInternalServerError, err)
		} else {
			fmt.Println(res.All(ctx, &urls))
			c.JSON(http.StatusOK, gin.H{
				"count": count,
				"date":  urls,
			})
			//c.JSON(http.StatusOK, urls)
		}

	}
}
