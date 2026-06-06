package controllers

import (
	"net/http"
	"strconv"

	"event-racing/src/models"
	"event-racing/src/services"

	"github.com/gin-gonic/gin"
)

func ListVenues(c *gin.Context) {
	venues, err := services.GetAllVenues()
	if err != nil {
		venues = []models.Venue{}
	}
	data := baseData(c)
	data["Title"] = "场地管理"
	data["Venues"] = venues
	c.HTML(http.StatusOK, "venues/list.html", data)
}

func GetVenue(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	venue, err := services.GetVenueByID(id)
	if err != nil {
		c.Redirect(http.StatusFound, "/venues")
		return
	}
	schedules, _ := services.GetSchedulesByVenue(id)
	data := baseData(c)
	data["Title"] = "场地详情"
	data["Venue"] = venue
	data["Schedules"] = schedules
	c.HTML(http.StatusOK, "venues/detail.html", data)
}

func ShowVenueSchedule(c *gin.Context) {
	venues, _ := services.GetAllVenues()
	data := baseData(c)
	data["Title"] = "场地调度"
	data["Venues"] = venues
	c.HTML(http.StatusOK, "venues/schedule.html", data)
}

func CreateVenue(c *gin.Context) {
	capacity, _ := strconv.Atoi(c.PostForm("capacity"))
	venue := &models.Venue{
		Name:     c.PostForm("name"),
		Location: c.PostForm("location"),
		Capacity: capacity,
		Status:   c.PostForm("status"),
	}
	if venue.Status == "" {
		venue.Status = "available"
	}
	if err := services.CreateVenue(venue); err != nil {
		c.Redirect(http.StatusFound, "/venues")
		return
	}
	c.Redirect(http.StatusFound, "/venues")
}

func UpdateVenue(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	capacity, _ := strconv.Atoi(c.PostForm("capacity"))
	venue := &models.Venue{
		ID:       id,
		Name:     c.PostForm("name"),
		Location: c.PostForm("location"),
		Capacity: capacity,
		Status:   c.PostForm("status"),
	}
	if err := services.UpdateVenue(venue); err != nil {
		c.Redirect(http.StatusFound, "/venues")
		return
	}
	c.Redirect(http.StatusFound, "/venues/"+strconv.FormatInt(id, 10))
}

func DeleteVenue(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	services.DeleteVenue(id)
	c.Redirect(http.StatusFound, "/venues")
}
