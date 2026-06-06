package controllers

import (
	"net/http"
	"strconv"

	"event-racing/src/models"
	"event-racing/src/services"

	"github.com/gin-gonic/gin"
)

func ListAwards(c *gin.Context) {
	awards, err := services.GetAllAwards()
	if err != nil {
		awards = []models.Award{}
	}
	data := baseData(c)
	data["Title"] = "获奖档案"
	data["Awards"] = awards
	c.HTML(http.StatusOK, "awards/list.html", data)
}

func GetAward(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	award, err := services.GetAwardByID(id)
	if err != nil {
		c.Redirect(http.StatusFound, "/awards")
		return
	}
	data := baseData(c)
	data["Title"] = "获奖详情"
	data["Award"] = award
	c.HTML(http.StatusOK, "awards/detail.html", data)
}

func CreateAward(c *gin.Context) {
	athleteID, _ := strconv.ParseInt(c.PostForm("athlete_id"), 10, 64)
	award := &models.Award{
		AthleteID:       athleteID,
		Event:           c.PostForm("event"),
		MedalType:       c.PostForm("medal_type"),
		CompetitionName: c.PostForm("competition_name"),
	}
	if err := services.CreateAward(award); err != nil {
		c.Redirect(http.StatusFound, "/awards")
		return
	}
	c.Redirect(http.StatusFound, "/awards")
}

func DeleteAward(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	services.DeleteAward(id)
	c.Redirect(http.StatusFound, "/awards")
}
