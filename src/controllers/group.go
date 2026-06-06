package controllers

import (
	"net/http"
	"strconv"

	"event-racing/src/models"
	"event-racing/src/services"

	"github.com/gin-gonic/gin"
)

func ListGroups(c *gin.Context) {
	groups, err := services.GetAllGroups()
	if err != nil {
		groups = []models.Group{}
	}
	data := baseData(c)
	data["Title"] = "分组编排"
	data["Groups"] = groups
	c.HTML(http.StatusOK, "groups/list.html", data)
}

func GetGroup(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	group, err := services.GetGroupByID(id)
	if err != nil {
		c.Redirect(http.StatusFound, "/groups")
		return
	}
	data := baseData(c)
	data["Title"] = "分组详情"
	data["Group"] = group
	c.HTML(http.StatusOK, "groups/detail.html", data)
}

func ShowArrangeGroup(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	group, err := services.GetGroupByID(id)
	if err != nil {
		c.Redirect(http.StatusFound, "/groups")
		return
	}
	athletes, _ := services.GetAthletesByEvent(group.Event)
	if athletes == nil {
		athletes = []models.Athlete{}
	}
	data := baseData(c)
	data["Title"] = "编排分组"
	data["Group"] = group
	data["AvailableAthletes"] = athletes
	c.HTML(http.StatusOK, "groups/arrange.html", data)
}

func ArrangeGroup(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	athleteIDs := c.PostFormArray("athlete_ids")
	var ids []int64
	for _, aid := range athleteIDs {
		if aidInt, err := strconv.ParseInt(aid, 10, 64); err == nil {
			ids = append(ids, aidInt)
		}
	}
	if err := services.ArrangeAthletes(id, ids); err != nil {
		data := baseData(c)
		data["Title"] = "编排分组"
		data["Error"] = "编排失败：" + err.Error()
		c.HTML(http.StatusBadRequest, "groups/arrange.html", data)
		return
	}
	c.Redirect(http.StatusFound, "/groups/"+strconv.FormatInt(id, 10))
}

func ShowCreateGroup(c *gin.Context) {
	data := baseData(c)
	data["Title"] = "创建分组"
	c.HTML(http.StatusOK, "groups/list.html", data)
}

func CreateGroup(c *gin.Context) {
	group := &models.Group{
		Name:   c.PostForm("name"),
		Event:  c.PostForm("event"),
		Status: "pending",
	}
	if err := services.CreateGroup(group); err != nil {
		c.Redirect(http.StatusFound, "/groups")
		return
	}
	c.Redirect(http.StatusFound, "/groups")
}

func DeleteGroup(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	services.DeleteGroup(id)
	c.Redirect(http.StatusFound, "/groups")
}
