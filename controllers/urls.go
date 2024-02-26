package controllers

import (
	"context"
	"github.com/MrMohebi/golang-gin-boilerplate.git/common"
	"github.com/MrMohebi/golang-gin-boilerplate.git/models"
	"github.com/gin-gonic/gin"
	"math/rand"
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

		hashedUrl := common.ConvertUrlToBase48(reqBody.OriginalUrl)

		url := models.Url{
			Id:          rand.Int(),
			Title:       reqBody.Title,
			Hash:        hashedUrl,
			OriginalUrl: reqBody.OriginalUrl,
			CreatedAt:   now.Format(time.DateTime),
			UpdatedAt:   now.Format(time.DateTime),
		}

		_, err = urlsColl.InsertOne(context.TODO(), url)
		common.IsErr(err, true)

		url.Hash = url.Hash[:5]

		ctx.JSON(http.StatusOK, url)
	}
}
