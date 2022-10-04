package act

import (
	h "ditto/ctrl"
	"ditto/lib/datetime"
	"ditto/lib/format"
	"ditto/model/act"
	"ditto/mw"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var sw *act.StopWatch

func Route(e *gin.Engine) {
	act := e.Group("/act").Use(mw.Auth)
	{
		act.GET("/", index)
		act.Any("/create", create)
		act.Any("/watch/start", start)
		act.Any("/watch/started", started)
		act.POST("/watch/stop", stop)
		act.GET("/delete", delete)
		act.DELETE("/delete", delete)
	}

	api := e.Group("/api/act")
	{
		api.GET("/", index)
		api.GET("/stopwatch", stopwatch)
		api.DELETE("/delete", delete).Use(mw.Auth)
	}
}

func create(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.HTML(http.StatusOK, "act/create", gin.H{
			"games": h.GameService.ByPlaying(),
		})

	case "POST":
		t := act.Type(c.PostForm("type"))
		date := datetime.FormatDate(c.PostForm("date"), datetime.DEFAULT)
		dur, _ := strconv.Atoi(c.PostForm("duration"))
		gId := format.ToObjId(c.PostForm("game_id"))
		h.ActService.Create(act.Act{
			Type:     t,
			Date:     date,
			Duration: dur,
			GameID:   gId,
		})
		c.Redirect(http.StatusSeeOther, "/act")
	}
}

func delete(c *gin.Context) {

	id := c.PostForm("id")
	err := h.ActService.Delete(id)
	if err != nil {
		c.JSON(404, gin.H{"msg": "delete failed!"})
		return
	}
	c.JSON(200, gin.H{"msg": "delete successed!"})
}

func index(c *gin.Context) {
	// Date format is yyyy-MM-dd
	date := c.Query("date")

	if len(date) == 0 {
		date = datetime.Today(datetime.DEFAULT)
	} else if len(date) > 6 {
		date = datetime.FormatDate(date, datetime.DEFAULT)
	}

	dailies, _ := h.ActService.ByDate(datetime.FormatDate(date, datetime.DEFAULT))
	monthlies, _ := h.ActService.ByMonth(datetime.FormatDate(date, datetime.DEFAULT))

	daySum := h.ActService.DaySum(datetime.FormatDate(date, datetime.DEFAULT))
	monthSum := h.ActService.MonthSum(date[0:6])

	data := gin.H{
		"date":                 datetime.FormatDate(date, datetime.HYPHEN),
		"daily_acts":           dailies,
		"day_sum":              daySum,
		"monthly_acts":         monthlies,
		"month_sum":            monthSum,
		"playing_games":        h.GameService.ByPlaying(),
		"is_stopwatch_started": sw != nil,
	}

	h.RESP(c, http.StatusOK, "act/index", data)
}

func start(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		// GET is from game index page
		sw = act.NewStopWatch()
		typ := c.Query("type")
		gid := c.Query("game_id")
		g := h.GameService.ByID(gid)
		sw.Start(typ, gid, g.Title)
	case "POST":
		// POST is from act index page
		typ := c.PostForm("type")
		if sw == nil && typ != "Recover" {
			sw = act.NewStopWatch()
			gid := c.PostForm("game_id")
			gtl := c.PostForm("game_title")
			sw.Start(typ, gid, gtl)
		}
	}
	c.Redirect(http.StatusSeeOther, "/act/watch/started")
}

func started(c *gin.Context) {
	c.HTML(http.StatusOK, "watch/started", gin.H{
		"watch": sw,
	})
	// c.Redirect(http.StatusSeeOther, "/watch/started")
}

func stop(c *gin.Context) {
	sw.Stop()
	if sw.Duration > 5 {
		h.ActService.Create(act.Act{
			StartTime: sw.StartTime,
			EndTime:   sw.EndTime,
			Date:      datetime.Today(datetime.DEFAULT),
			Duration:  sw.Duration,
			GameID:    format.ToObjId(sw.GameID),
			Type:      sw.Type,
		})
	}
	sw = nil
	c.Redirect(http.StatusSeeOther, "/act")
}

func stopwatch(c *gin.Context) {
	var watch act.StopWatch
	if sw != nil {
		watch = *sw
	}
	h.RESP(c, http.StatusOK, "act/index", watch)
}
