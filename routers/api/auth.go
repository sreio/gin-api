package api

import (
	"gin-api/models"
	"gin-api/pkg/e"
	"gin-api/pkg/logging"
	"gin-api/pkg/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}


// @Summary jwt验证
// @Produce  json
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {json} gin.H
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	v := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := v.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				code = e.SUCCESS
				data["token"] = token
			}
		} else {
			code = e.ERROR_USER_EXIST
		}
	} else {
		for _, err := range v.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
