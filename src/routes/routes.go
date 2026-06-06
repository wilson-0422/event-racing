package routes

import (
	"event-racing/src/controllers"
	"event-racing/src/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(middleware.SetUserContext())

	r.GET("/", controllers.ShowIndex)
	r.GET("/dashboard", middleware.AuthRequired(), controllers.ShowDashboard)

	auth := r.Group("/auth")
	{
		auth.GET("/login", controllers.ShowLogin)
		auth.POST("/login", controllers.Login)
		auth.GET("/register", controllers.ShowRegister)
		auth.POST("/register", controllers.Register)
		auth.GET("/logout", controllers.Logout)
	}

	athletes := r.Group("/athletes")
	athletes.Use(middleware.AuthRequired())
	{
		athletes.GET("", controllers.ListAthletes)
		athletes.GET("/create", controllers.ShowCreateAthlete)
		athletes.POST("", controllers.CreateAthlete)
		athletes.GET("/:id", controllers.GetAthlete)
		athletes.GET("/:id/edit", controllers.ShowEditAthlete)
		athletes.POST("/:id", controllers.UpdateAthlete)
		athletes.POST("/:id/delete", controllers.DeleteAthlete)
	}

	groups := r.Group("/groups")
	groups.Use(middleware.AuthRequired())
	{
		groups.GET("", controllers.ListGroups)
		groups.POST("", controllers.CreateGroup)
		groups.GET("/:id", controllers.GetGroup)
		groups.GET("/:id/arrange", controllers.ShowArrangeGroup)
		groups.POST("/:id/arrange", controllers.ArrangeGroup)
		groups.POST("/:id/delete", controllers.DeleteGroup)
	}

	schedules := r.Group("/schedules")
	schedules.Use(middleware.AuthRequired())
	{
		schedules.GET("", controllers.ListSchedules)
		schedules.GET("/create", controllers.ShowCreateSchedule)
		schedules.POST("", controllers.CreateSchedule)
		schedules.GET("/:id", controllers.GetSchedule)
		schedules.POST("/:id/status", controllers.UpdateScheduleStatus)
		schedules.POST("/:id/delete", controllers.DeleteSchedule)
	}

	scores := r.Group("/scores")
	scores.Use(middleware.AuthRequired())
	{
		scores.GET("", controllers.ListScores)
		scores.GET("/entry", controllers.ShowScoreEntry)
		scores.POST("", controllers.CreateScore)
		scores.GET("/:id", controllers.GetScore)
		scores.POST("/:id/delete", controllers.DeleteScore)
	}

	awards := r.Group("/awards")
	awards.Use(middleware.AuthRequired())
	{
		awards.GET("", controllers.ListAwards)
		awards.POST("", controllers.CreateAward)
		awards.GET("/:id", controllers.GetAward)
		awards.POST("/:id/delete", controllers.DeleteAward)
	}

	venues := r.Group("/venues")
	venues.Use(middleware.AuthRequired())
	{
		venues.GET("", controllers.ListVenues)
		venues.POST("", controllers.CreateVenue)
		venues.GET("/schedule", controllers.ShowVenueSchedule)
		venues.GET("/:id", controllers.GetVenue)
		venues.POST("/:id", controllers.UpdateVenue)
		venues.POST("/:id/delete", controllers.DeleteVenue)
	}
}
