package act

import (
	h "ditto/ctrl"
	"ditto/lib/datetime"
	"ditto/lib/format"
	"ditto/model/act"
	"ditto/mw"
	"net/http"

	"github.com/gin-gonic/gin"
)

var sw *act.StopWatch

func Route(e *gin.Engine) {
	auth := e.Group("/act").Use(mw.Auth)
	{
		// auth.GET("/", index)
		// auth.POST("/create", create)
		auth.Any("/watch/start", start)
		auth.Any("/watch/started", started)
		auth.POST("/watch/stop", stop)
		auth.GET("/delete", delete)
		auth.DELETE("/delete", delete)
		auth.GET("/stopwatch", stopwatch)
	}

	anon := e.Group("/act")
	{
		anon.GET("/", index)
		anon.POST("/create", create)
	}
}

func create(c *gin.Context) {
	type actJson struct {
		Type     string `json:"type"`
		Date     string `json:"date"`
		Duration int    `json:"duration"`
		GameId   string `json:"game_id"`
	}

	var json actJson
	if err := c.BindJSON(&json); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	t := act.Type(json.Type)
	date := datetime.FormatDate(json.Date, datetime.DEFAULT)
	dur := json.Duration
	gId := format.ToObjId(json.GameId)
	h.ActService.Create(act.Act{
		Type:     t,
		Date:     date,
		Duration: dur,
		GameID:   gId,
	})
	c.JSON(http.StatusAccepted, "")
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
	date := c.Query("date") // YYYYMMDD

	if len(date) == 0 {
		date = datetime.Today(datetime.DEFAULT)
	} else if len(date) > 6 {
		date = datetime.FormatDate(date, datetime.DEFAULT)
	}

	dayDetails, _ := h.ActService.ByDate(date)
	daySummary := h.ActService.DaySum(date)
	monDetails, _ := h.ActService.ByMonth(datetime.FormatDate(date, datetime.DEFAULT))
	monSummary := h.ActService.MonthSum(date[0:6])

	data := gin.H{
		"date":          datetime.FormatDate(date, datetime.HYPHEN),
		"day_details":   dayDetails,
		"day_summary":   daySummary,
		"month_details": monDetails,
		"month_summary": monSummary,
		"playing_games": h.GameService.ByPlaying(),
	}
	c.JSON(200, data)
}

func start(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		sw = act.NewStopWatch()
		typ := "Gaming"
		gid := c.Query("id")
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
