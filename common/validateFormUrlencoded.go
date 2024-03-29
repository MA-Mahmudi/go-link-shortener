package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// ValidBindForm bind request body(Form urlencode) and validate fields with Validator-v10. If err returns json response to gin
func ValidBindForm(c *gin.Context, reqBody interface{}) bool {
	var validate = validator.New()
	if err := c.Bind(reqBody); err != nil {
		c.JSON(http.StatusBadRequest, http.StatusBadRequest)
		return false
	}
	if validationErr := validate.Struct(reqBody); validationErr != nil {
		c.JSON(http.StatusBadRequest, validationErr.Error())
		return false
	}
	return true
}

func IsValidFormUrlencoded(reqBody *struct{}) bool {
	var validate = validator.New()
	if validationErr := validate.Struct(&reqBody); validationErr != nil {
		return false
	}
	return true
}

func ValidRawJson(c *gin.Context, reqBody interface{}) bool {
	var validate = validator.New()
	err := c.BindJSON(reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return false
	}
	validationErr := validate.Struct(reqBody)
	if validationErr != nil {
		fmt.Println(validationErr.Error())
		c.JSON(http.StatusUnprocessableEntity, validationErr.Error())
		return false
	}
	return true
}
