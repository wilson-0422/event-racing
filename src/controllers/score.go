package controllers

import (
	"net/http"
	"strconv"

	"event-racing/src/models"
	"event-racing/src/services"

	"github.com/gin-gonic/gin"
)

func ListScores(c *gin.Context) {
	scores, err := services.GetAllScores()
	if err != nil {
		scores = []models.Score{}
	}
	data := baseData(c)
	data["Title"] = "成绩管理"
	data["Scores"] = scores
	c.HTML(http.StatusOK, "scores/list.html", data)
}

func GetScore(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	score, err := services.GetScoreByID(id)
	if err != nil {
		c.Redirect(http.StatusFound, "/scores")
		return
	}
	data := baseData(c)
	data["Title"] = "成绩详情"
	data["Score"] = score
	c.HTML(http.StatusOK, "scores/detail.html", data)
}

func ShowScoreEntry(c *gin.Context) {
	schedules, _ := services.GetAllSchedules()
	athletes, _ := services.GetAllAthletes()
	data := baseData(c)
	data["Title"] = "成绩录入"
	data["Schedules"] = schedules
	data["Athletes"] = athletes
	c.HTML(http.StatusOK, "scores/entry.html", data)
}

func CreateScore(c *gin.Context) {
	athleteID, _ := strconv.ParseInt(c.PostForm("athlete_id"), 10, 64)
	scheduleID, _ := strconv.ParseInt(c.PostForm("schedule_id"), 10, 64)
	rank, _ := strconv.Atoi(c.PostForm("rank"))
	score := &models.Score{
		AthleteID:  athleteID,
		ScheduleID: scheduleID,
		Score:      c.PostForm("score"),
		Rank:       rank,
		Remark:     c.PostForm("remark"),
	}
	if err := services.CreateScore(score); err != nil {
		schedules, _ := services.GetAllSchedules()
		athletes, _ := services.GetAllAthletes()
		data := baseData(c)
		data["Title"] = "成绩录入"
		data["Schedules"] = schedules
		data["Athletes"] = athletes
		data["Error"] = "录入失败：" + err.Error()
		c.HTML(http.StatusBadRequest, "scores/entry.html", data)
		return
	}
	c.Redirect(http.StatusFound, "/scores")
}

func DeleteScore(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	services.DeleteScore(id)
	c.Redirect(http.StatusFound, "/scores")
}
