package game

import (
	h "ditto/ctrl"
	"ditto/db/mongo"
	"ditto/lib/format"
	"ditto/lib/path"
	"ditto/model/game"
	"ditto/mw"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	PAGE_LIMIT = 9
)

func Route(e *gin.Engine) {
	auth := e.Group("/game").Use(mw.Auth)
	{
		auth.GET("/", index)
		auth.Any("/create", create)
		auth.Any("/update", update)
		auth.Any("/delete", delete)
		auth.POST("/update_status", updateStatus)

		auth.GET("/counts", counts)
		auth.GET("/status/:status/:platform/:page", status)
	}

	// anon := e.Group("/api/game")
	// {
	// 	anon.GET("/counts", counts)
	// 	anon.GET("/", index)
	// 	anon.GET("/status/:status/:platform/:page", status)
	// }
}

func updateStatus(c *gin.Context) {
	// json, _ := ioutil.ReadAll(c.Request.Body)

	gameId := c.PostForm("id")
	newStatus := c.PostForm("newStatus")

	targetGame := h.GameService.ByID(gameId)
	targetGame.Status = game.Status(newStatus)

	h.GameService.Update(targetGame)
}

func counts(c *gin.Context) {
	playedCnt, playingCnt, blockingCnt := h.GameService.Counts()
	c.JSON(http.StatusOK, gin.H{
		"played_cnt":   playedCnt,
		"playing_cnt":  playingCnt,
		"blocking_cnt": blockingCnt,
	})
}

func create(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.HTML(http.StatusOK, "game/create", gin.H{
			"developers": h.IncService.Developers(),
			"publishers": h.IncService.Publishers(),
			"genres":     game.Genres(),
			"platforms":  game.Platforms(),
		})
	case "POST":
		title := c.PostForm("title")
		developerId := c.PostForm("developer_id")
		publisherId := c.PostForm("publisher_id")
		genre := c.PostForm("genre")
		platform := game.Platform(c.PostForm("platform"))

		h.GameService.Create(game.Game{
			Title:       title,
			Status:      game.PLAYING,
			Genre:       genre,
			Platform:    platform,
			DeveloperID: format.ObjId(developerId),
			PublisherID: format.ObjId(publisherId),
		})
		c.Redirect(http.StatusSeeOther, "/game")
	}
}

func index(c *gin.Context) {
	status := game.Status(c.Query("status"))
	data := gin.H{
		"details": h.GameService.ByStatus(status),
	}

	c.JSON(http.StatusOK, data)
}

func update(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		id := c.Query("id")
		g := h.GameService.ByID(id)

		c.HTML(http.StatusOK, "game/update", gin.H{
			"game":           g,
			"developers":     h.IncService.Developers(),
			"publishers":     h.IncService.Publishers(),
			"statuses":       game.Statuses(),
			"platforms":      game.Platforms(),
			"genres":         game.Genres(),
			"play_time_hour": g.PlayTime / 60,
			"play_time_min":  g.PlayTime % 60,
		})

	case "POST":
		gId := c.PostForm("id")
		dId := c.PostForm("developer_id")
		pId := c.PostForm("publisher_id")
		pth, _ := strconv.Atoi(c.PostForm("play_time_hour"))
		ptm, _ := strconv.Atoi(c.PostForm("play_time_min"))
		pt := pth*60 + ptm
		ranking, _ := strconv.Atoi(c.PostForm("ranking"))
		rating := c.PostForm("rating")

		g := h.GameService.ByID(gId)
		g.Title = c.PostForm("title")
		g.DeveloperID = format.ObjId(dId)
		g.PublisherID = format.ObjId(pId)
		g.Status = game.Status(c.PostForm("status"))
		g.PlayTime = pt
		g.Genre = c.PostForm("genre")
		g.Platform = game.Platform(c.PostForm("platform"))
		g.Ranking = ranking
		g.Rating = rating

		file, err := c.FormFile("cover")
		// Upload image to assets
		if file != nil && err == nil {
			fn := gId + ".webp"
			root := path.Root()
			path := root + "/../assets/img/games/" + fn
			c.SaveUploadedFile(file, path)
		}

		// Uplaod image to Firebase storage
		// if err == nil {
		// 	firebase.Upload(file, gId+".webp")
		// }

		h.GameService.Update(g)
		c.Redirect(http.StatusSeeOther, "/game")
	}
}

func delete(c *gin.Context) {
	id := c.Query("id")

	// Delete from db.act
	acts, err := h.ActService.ByGame(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, err)
		return
	}

	if len(acts) != 0 {
		for _, v := range acts {
			mongo.DeleteByID(mongo.Acts, v.ID)
		}
	}

	// Delete from db.game
	err = h.GameService.Delete(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, err)
		return
	}

	// Delete from assets
	file := id + ".webp"
	root := path.Root()
	path := root + "/../assets/img/games/" + file
	os.Remove(path)

	c.Redirect(http.StatusSeeOther, "/game")
}

func status(c *gin.Context) {
	status := game.Status(c.Param("status"))
	platform := game.Platform(c.Param("platform"))

	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		page = 1
	}

	details, totalPages := h.GameService.PageByStatus(status, platform, page, PAGE_LIMIT)
	data := gin.H{
		"details":     details,
		"total_pages": totalPages,
	}

	c.JSON(http.StatusOK, data)
}
