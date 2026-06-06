package controllers

import (
	"net/http"
	"strconv"

	"event-racing/src/models"
	"event-racing/src/services"

	"github.com/gin-gonic/gin"
)

func ListSchedules(c *gin.Context) {
	schedules, err := services.GetAllSchedules()
	if err != nil {
		schedules = []models.Schedule{}
	}
	data := baseData(c)
	data["Title"] = "赛程排班"
	data["Schedules"] = schedules
	c.HTML(http.StatusOK, "schedules/list.html", data)
}

func GetSchedule(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	schedule, err := services.GetScheduleByID(id)
	if err != nil {
		c.Redirect(http.StatusFound, "/schedules")
		return
	}
	scores, _ := services.GetScoresBySchedule(id)
	data := baseData(c)
	data["Title"] = "赛程详情"
	data["Schedule"] = schedule
	data["Scores"] = scores
	c.HTML(http.StatusOK, "schedules/detail.html", data)
}

func ShowCreateSchedule(c *gin.Context) {
	groups, _ := services.GetAllGroups()
	venues, _ := services.GetAllVenues()
	data := baseData(c)
	data["Title"] = "创建赛程"
	data["Groups"] = groups
	data["Venues"] = venues
	c.HTML(http.StatusOK, "schedules/create.html", data)
}

func CreateSchedule(c *gin.Context) {
	groupID, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	venueID, _ := strconv.ParseInt(c.PostForm("venue_id"), 10, 64)
	schedule := &models.Schedule{
		GroupID:   groupID,
		VenueID:   venueID,
		StartTime: c.PostForm("start_time"),
		EndTime:   c.PostForm("end_time"),
		Status:    "scheduled",
	}
	if err := services.CreateSchedule(schedule); err != nil {
		groups, _ := services.GetAllGroups()
		venues, _ := services.GetAllVenues()
		data := baseData(c)
		data["Title"] = "创建赛程"
		data["Groups"] = groups
		data["Venues"] = venues
		data["Error"] = "创建失败：" + err.Error()
		c.HTML(http.StatusBadRequest, "schedules/create.html", data)
		return
	}
	c.Redirect(http.StatusFound, "/schedules")
}

func UpdateScheduleStatus(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	status := c.PostForm("status")
	schedule, err := services.GetScheduleByID(id)
	if err != nil {
		c.Redirect(http.StatusFound, "/schedules")
		return
	}
	schedule.Status = status
	services.UpdateSchedule(schedule)
	c.Redirect(http.StatusFound, "/schedules/"+strconv.FormatInt(id, 10))
}

func DeleteSchedule(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	services.DeleteSchedule(id)
	c.Redirect(http.StatusFound, "/schedules")
}
