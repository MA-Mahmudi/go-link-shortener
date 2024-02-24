package contorolers

import (
	"context"
	"encoding/json"
	"github.com/MrMohebi/golang-gin-boilerplate.git/common"
	"github.com/MrMohebi/golang-gin-boilerplate.git/models"
	"github.com/gin-gonic/gin"
	"io"
)

func CreateUrl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		coll := models.UrlsCollection

		url, err := io.ReadAll(ctx.Request.Body)

		res, err := coll.InsertOne(context.TODO(), url)
		common.IsErr(err, true)

		println(json.Marshal(res.InsertedID))

	}

}
