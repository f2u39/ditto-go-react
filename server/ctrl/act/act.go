package act

import (
	h "ditto/ctrl"
	"ditto/lib/datetime"
	"ditto/lib/format"
	"ditto/model/act"
	"ditto/mw"
	"net/http"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var sw *act.StopWatch

func Route(e *gin.Engine) {
	// Use React as frontend
	e.Use(static.Serve("/act", static.LocalFile("./web", true)))

	auth := e.Group("/api/act").Use(mw.Auth)
	{
		auth.Any("/watch/start", start)
		auth.POST("/watch/stop", stop)
		auth.POST("/watch/teminate", terminate)
		auth.POST("/create", create)
		auth.GET("/delete", delete)
		auth.DELETE("/delete", delete)
		auth.GET("/stopwatch", stopwatch)
	}

	api := e.Group("/api/act")
	{
		api.GET("/", index)
	}
}

func create(c *gin.Context) {
	type actJson struct {
		Type     string `json:"type"`
		Date     string `json:"date"`
		Duration string `json:"duration"`
		GameId   string `json:"gameId"`
	}

	var json actJson
	if err := c.BindJSON(&json); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	typ := act.Type(json.Type)
	date := datetime.FormatDate(json.Date, datetime.DEFAULT)
	dur, _ := strconv.Atoi(json.Duration)
	gId := format.ToObjID(json.GameId)

	if typ == act.GAMING {
		if len(gId) == 0 {
			c.JSON(http.StatusBadRequest, "")
			return
		}
	}

	h.ActService.Create(act.Act{
		Type:     typ,
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
		"stopwatch":     sw,
	}

	c.JSON(http.StatusOK, data)
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
		type swJson struct {
			Type   string `json:"type"`
			GameId string `json:"gameId"`
		}

		var json swJson
		if err := c.BindJSON(&json); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		typ := json.Type
		gId := json.GameId

		if act.Type(typ) == act.GAMING {
			if len(gId) == 0 {
				c.JSON(http.StatusBadRequest, "Illeagal Game Id")
				return
			}
		}

		if sw == nil && typ != "Recover" {
			sw = act.NewStopWatch()
			id := gId
			title := h.GameService.ByID(gId).Title
			sw.Start(typ, id, title)
		}
		data := gin.H{
			"stop_watch": sw,
		}
		c.JSON(http.StatusOK, data)
	}
}

func stop(c *gin.Context) {
	sw.Stop()

	// At least 5 minutes
	if sw.Duration > 5 {
		h.ActService.Create(act.Act{
			StartTime: sw.StartTime,
			EndTime:   sw.EndTime,
			Date:      datetime.Today(datetime.DEFAULT),
			Duration:  sw.Duration,
			GameID:    format.ToObjID(sw.GameID),
			Type:      sw.Type,
		})
	}

	sw.Reset()
	c.JSON(http.StatusOK, nil)
}

func terminate(c *gin.Context) {
	sw.Reset()
}

func stopwatch(c *gin.Context) {
	var watch act.StopWatch
	if sw != nil {
		watch = *sw
	}
	c.JSON(http.StatusOK, watch)
}
