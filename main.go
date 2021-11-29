package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type day struct {
	Day   string
	Text  string
	Image string
	Class string
}

type week struct {
	Week []day
}

func getDays() []week {

	month := []week{
		{
			[]day{{
				Class: "empty",
			}, {
				Class: "standard",
				Text:  "event",
				Day:   "1",
			}, {
				Class: "important",
				Text:  "birthday",
				Day:   "2",
			},
			}},
		{[]day{}},
	}

	return month
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/css", "./templates/css")

	r.GET("/", func(c *gin.Context) {
		month := 11
		year := 2021
		title := "Printable Calendar"
		location, _ := time.LoadLocation("America/New_York")

		startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, location)
		daysInMonth := startDate.AddDate(0, 1, -1).Day()
		weeks := getDays()

		calendarLength := "default"
		if month == 2 && daysInMonth == 28 && startDate.Weekday() == 0 {
			calendarLength = "short"
		} else if (daysInMonth >= 30 && startDate.Weekday() == 7) || (daysInMonth >= 31 && startDate.Weekday() == 6) {
			calendarLength = "long"
		}
		c.HTML(http.StatusOK, "calendar.tmpl.html", gin.H{
			"Month":  startDate.Month(),
			"Year":   startDate.Year(),
			"Title":  title,
			"Length": calendarLength,
			"Weeks":  weeks,
		})
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Serve in 0.0.0.0:8080
	r.Run(":8080")
}
