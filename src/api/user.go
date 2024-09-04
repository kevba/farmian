package api

import (
	"farmian/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	login := router.Group("login")
	login.POST("", loginUser)

	user := router.Group("user")
	user.Use(Auth())
	user.GET("", getUser)

}

func loginUser(c *gin.Context) {
	type loginBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	body := &loginBody{}
	_ = c.BindJSON(body)
	user := findUser(body.Username)

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"ok": false})
		return
	}

	if user.Password != body.Password {
		c.JSON(http.StatusNotFound, gin.H{"ok": false})
		return
	}

	token, _ := createToken(user)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func getUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success!"})
}

func findUser(name string) *models.User {
	users := models.GetUsers()

	for _, u := range users {
		if u.Username == name {
			return u
		}
	}
	return nil
}
