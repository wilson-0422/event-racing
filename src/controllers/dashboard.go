package controllers

import (
	"net/http"

	"event-racing/src/models"
	"event-racing/src/services"

	"github.com/gin-gonic/gin"
)

func ShowDashboard(c *gin.Context) {
	athleteCount, _ := services.CountAthletes()
	groupCount, _ := services.CountGroups()
	scheduleCount, _ := services.CountSchedules()
	awardCount, _ := services.CountAwards()
	venueCount, _ := services.CountVenues()
	athletes, _ := services.GetAllAthletes()
	schedules, _ := services.GetAllSchedules()
	awards, _ := services.GetAllAwards()

	if len(athletes) > 5 {
		athletes = athletes[:5]
	}
	if len(schedules) > 5 {
		schedules = schedules[:5]
	}
	if len(awards) > 5 {
		awards = awards[:5]
	}

	data := baseData(c)
	data["Title"] = "系统概览"
	data["AthleteCount"] = athleteCount
	data["GroupCount"] = groupCount
	data["ScheduleCount"] = scheduleCount
	data["AwardCount"] = awardCount
	data["VenueCount"] = venueCount
	data["RecentAthletes"] = athletes
	data["UpcomingSchedules"] = schedules
	data["RecentAwards"] = awards
	c.HTML(http.StatusOK, "dashboard/overview.html", data)
}

func ShowIndex(c *gin.Context) {
	athleteCount, _ := services.CountAthletes()
	groupCount, _ := services.CountGroups()
	scheduleCount, _ := services.CountSchedules()
	awardCount, _ := services.CountAwards()
	venueCount, _ := services.CountVenues()

	data := baseData(c)
	data["Title"] = "首页"
	data["AthleteCount"] = athleteCount
	data["GroupCount"] = groupCount
	data["ScheduleCount"] = scheduleCount
	data["AwardCount"] = awardCount
	data["VenueCount"] = venueCount
	c.HTML(http.StatusOK, "index.html", data)
}

func limitAthletes(slice []models.Athlete, n int) []models.Athlete {
	if len(slice) > n {
		return slice[:n]
	}
	return slice
}

func limitSchedules(slice []models.Schedule, n int) []models.Schedule {
	if len(slice) > n {
		return slice[:n]
	}
	return slice
}

func limitAwards(slice []models.Award, n int) []models.Award {
	if len(slice) > n {
		return slice[:n]
	}
	return slice
}
