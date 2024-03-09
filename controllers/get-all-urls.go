package controllers

import (
	"context"
	"fmt"
	"github.com/MrMohebi/golang-gin-boilerplate.git/common"
	"github.com/MrMohebi/golang-gin-boilerplate.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

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
