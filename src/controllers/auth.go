package controllers

import (
	"net/http"
	"strconv"

	"event-racing/src/config"
	"event-racing/src/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/login.html", gin.H{
		"Title": "登录",
	})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.HTML(http.StatusBadRequest, "auth/login.html", gin.H{
			"Title": "登录",
			"Error": "用户名和密码不能为空",
		})
		return
	}
	user, err := services.VerifyPassword(username, password)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "auth/login.html", gin.H{
			"Title": "登录",
			"Error": "用户名或密码错误",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Set("username", user.Username)
	session.Set("user_role", user.Role)
	session.Save()
	c.Redirect(http.StatusFound, "/")
}

func ShowRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/register.html", gin.H{
		"Title": "注册",
	})
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")
	if username == "" || password == "" {
		c.HTML(http.StatusBadRequest, "auth/register.html", gin.H{
			"Title": "注册",
			"Error": "用户名和密码不能为空",
		})
		return
	}
	if password != confirmPassword {
		c.HTML(http.StatusBadRequest, "auth/register.html", gin.H{
			"Title": "注册",
			"Error": "两次密码输入不一致",
		})
		return
	}
	_, err := services.CreateUser(username, password, "user")
	if err != nil {
		c.HTML(http.StatusBadRequest, "auth/register.html", gin.H{
			"Title": "注册",
			"Error": "用户名已存在",
		})
		return
	}
	c.Redirect(http.StatusFound, "/auth/login")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/auth/login")
}

func getUserFromSession(c *gin.Context) gin.H {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		return gin.H{}
	}
	id, _ := strconv.ParseInt(userID.(string), 10, 64)
	user, err := services.GetUserByID(id)
	if err != nil {
		return gin.H{}
	}
	return gin.H{
		"User": gin.H{
			"ID":       user.ID,
			"Username": user.Username,
			"Role":     user.Role,
		},
	}
}

func baseData(c *gin.Context) gin.H {
	data := getUserFromSession(c)
	data["AppName"] = config.AppName
	return data
}
