package controllers

import (
	"context"
	"fmt"
	"github.com/MrMohebi/golang-gin-boilerplate.git/common"
	"github.com/MrMohebi/golang-gin-boilerplate.git/constant"
	"github.com/MrMohebi/golang-gin-boilerplate.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func CreateUrl() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		now := time.Now()

		var reqBody models.CreateUrlStruct

		isValid := common.ValidRawJson(c, &reqBody)
		if isValid {
			customUrl := reqBody.CustomUrl
			shortUrl := ""
			if len(customUrl) == 0 {
				shortUrl = common.RandStr(5)
			} else {
				shortUrl = customUrl
			}

			url := models.Url{
				Id:          primitive.NewObjectID(),
				Title:       reqBody.Title,
				ShortLink:   shortUrl,
				OriginalUrl: reqBody.OriginalUrl,
				CreatedAt:   now.Format(time.DateTime),
				UpdatedAt:   now.Format(time.DateTime),
				Status:      constant.StatusActive,
				ExpireDate:  reqBody.ExpireDate,
			}
			_, err := models.UrlsCollection.InsertOne(ctx, url)
			if err != nil {
				common.IsErr(err, false)
				return
			}

			urlLog := models.UrlLog{
				Id:         primitive.NewObjectID(),
				UrlId:      url.Id,
				ClickCount: 0,
				Timestamp:  now.Unix(),
				IpAddress:  "host",
				IsBot:      0,
			}
			_, err = models.UrlsLogsCollection.InsertOne(ctx, urlLog)
			if err != nil {
				common.IsErr(err, false)
				return
			}

			if err == nil {
				c.JSON(http.StatusOK, url)
			}
		}
	}

}

func GetAllUrls() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var urls []models.Url

		count, _ := models.UrlsCollection.CountDocuments(ctx, bson.M{})
		res, err := models.UrlsCollection.Find(ctx, bson.M{})
		if err != nil {
			common.IsErr(err, false)
			c.JSON(http.StatusInternalServerError, err)
		} else {
			fmt.Println(res.All(ctx, &urls))
			c.JSON(http.StatusOK, gin.H{
				"count": count,
				"date":  urls,
			})
		}
	}
}
