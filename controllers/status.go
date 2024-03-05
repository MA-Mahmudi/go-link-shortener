package controllers

import (
	"context"
	"github.com/MrMohebi/golang-gin-boilerplate.git/common"
	"github.com/MrMohebi/golang-gin-boilerplate.git/constant"
	"github.com/MrMohebi/golang-gin-boilerplate.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type deleteReqBody struct {
	Id primitive.ObjectID `json:"id" validate:"required"`
}

func DeleteUrl() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		now := time.Now()

		var reqBody deleteReqBody

		if isValid := common.ValidRawJson(c, &reqBody); isValid {
			_, err := models.UrlsCollection.UpdateOne(ctx, bson.M{"_id": reqBody.Id}, bson.M{
				"$set": bson.M{
					"status":     constant.StatusDeleted,
					"updated_at": now.Format(time.DateTime),
					"deleted_at": now.Format(time.DateTime),
				},
			})
			if err != nil {
				common.IsErr(err, false)
				c.JSON(http.StatusNotFound, gin.H{
					"message": "Url not found.",
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"success": "ok",
					"message": "Url has been deleted.",
				})
			}
		}
	}
}

func IsUrlActive(c *gin.Context, id primitive.ObjectID) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	now := time.Now()

	var urlToCheck models.Url

	err := models.UrlsCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&urlToCheck)
	if err != nil {
		common.IsErr(err, false)
		c.JSON(http.StatusNotFound, gin.H{"message": "Url not found."})
		return false
	}
	if len(urlToCheck.ExpireDate) > 0 && urlToCheck.Status == constant.StatusActive {
		expireDate, _ := time.ParseInLocation(time.DateTime, urlToCheck.ExpireDate, time.Local)

		if now.After(expireDate) {
			_, err := models.UrlsCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"status": constant.StatusExpired}})
			if err != nil {
				common.IsErr(err, false)
				c.JSON(http.StatusInternalServerError, err)
				return false
			}
			c.JSON(http.StatusForbidden, gin.H{"message": "This Url is expired!"})
			return false
		} else {
			return true
		}
	}
	c.JSON(http.StatusForbidden, gin.H{"message": "This Url is NOT active!"})
	return false
}
