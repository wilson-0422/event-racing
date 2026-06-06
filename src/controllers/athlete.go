package controllers

import (
	"net/http"
	"strconv"

	"event-racing/src/models"
	"event-racing/src/services"

	"github.com/gin-gonic/gin"
)

func ListAthletes(c *gin.Context) {
	athletes, err := services.GetAllAthletes()
	if err != nil {
		athletes = []models.Athlete{}
	}
	data := baseData(c)
	data["Title"] = "运动员管理"
	data["Athletes"] = athletes
	c.HTML(http.StatusOK, "athletes/list.html", data)
}

func GetAthlete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	athlete, err := services.GetAthleteByID(id)
	if err != nil {
		c.Redirect(http.StatusFound, "/athletes")
		return
	}
	scores, _ := services.GetScoresByAthlete(id)
	awards, _ := services.GetAwardsByAthlete(id)
	data := baseData(c)
	data["Title"] = "运动员详情"
	data["Athlete"] = athlete
	data["Scores"] = scores
	data["Awards"] = awards
	c.HTML(http.StatusOK, "athletes/detail.html", data)
}

func ShowCreateAthlete(c *gin.Context) {
	data := baseData(c)
	data["Title"] = "运动员报名"
	c.HTML(http.StatusOK, "athletes/register.html", data)
}

func CreateAthlete(c *gin.Context) {
	age, _ := strconv.Atoi(c.PostForm("age"))
	athlete := &models.Athlete{
		Name:     c.PostForm("name"),
		Gender:   c.PostForm("gender"),
		Age:      age,
		Team:     c.PostForm("team"),
		Event:    c.PostForm("event"),
		Phone:    c.PostForm("phone"),
		IDNumber: c.PostForm("id_number"),
	}
	if err := services.CreateAthlete(athlete); err != nil {
		data := baseData(c)
		data["Title"] = "运动员报名"
		data["Error"] = "创建失败：" + err.Error()
		c.HTML(http.StatusBadRequest, "athletes/register.html", data)
		return
	}
	c.Redirect(http.StatusFound, "/athletes")
}

func ShowEditAthlete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	athlete, err := services.GetAthleteByID(id)
	if err != nil {
		c.Redirect(http.StatusFound, "/athletes")
		return
	}
	data := baseData(c)
	data["Title"] = "编辑运动员"
	data["Athlete"] = athlete
	c.HTML(http.StatusOK, "athletes/register.html", data)
}

func UpdateAthlete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	age, _ := strconv.Atoi(c.PostForm("age"))
	athlete := &models.Athlete{
		ID:       id,
		Name:     c.PostForm("name"),
		Gender:   c.PostForm("gender"),
		Age:      age,
		Team:     c.PostForm("team"),
		Event:    c.PostForm("event"),
		Phone:    c.PostForm("phone"),
		IDNumber: c.PostForm("id_number"),
	}
	if err := services.UpdateAthlete(athlete); err != nil {
		c.Redirect(http.StatusFound, "/athletes")
		return
	}
	c.Redirect(http.StatusFound, "/athletes/"+strconv.FormatInt(id, 10))
}

func DeleteAthlete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	services.DeleteAthlete(id)
	c.Redirect(http.StatusFound, "/athletes")
}
