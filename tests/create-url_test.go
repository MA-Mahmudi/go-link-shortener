package tests

import (
	"bytes"
	"encoding/json"
	"github.com/MrMohebi/golang-gin-boilerplate.git/common"
	"github.com/MrMohebi/golang-gin-boilerplate.git/controllers"
	"github.com/MrMohebi/golang-gin-boilerplate.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateUrl(t *testing.T) {
	r := gin.Default()
	r.GET("/url/create", controllers.CreateUrl())
	mockResponse := models.Url{
		Id:          primitive.ObjectID{},
		Title:       "Zoodex",
		ShortLink:   "mmd",
		OriginalUrl: "https://zooedex.ir",
		CreatedAt:   time.Now().Format(time.DateTime),
		UpdatedAt:   time.Now().Format(time.DateTime),
		DeletedAt:   "",
		ExpireDate:  "",
		Status:      1,
	}

	body := models.CreateUrlStruct{OriginalUrl: "https://zoodex.ir", Title: "Zoodex", CustomUrl: "mmd"}

	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest("GET", "/url/create", bytes.NewBuffer(jsonBody))
	if err != nil {
		common.IsErr(err, true)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	response, err := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(response))
	assert.Equal(t, http.StatusOK, w.Code)
}
