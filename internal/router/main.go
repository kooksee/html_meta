package router

import (
	"github.com/Masterminds/sprig"
	"github.com/gin-gonic/gin"
	"github.com/kooksee/html_meta/internal/config"
	"github.com/kooksee/html_meta/internal/services"
	"net/http"
	"time"
)

func App() *gin.Engine {
	cfg := config.DefaultConfig()
	cfg.Init()

	r := gin.Default()
	r.SetFuncMap(sprig.HtmlFuncMap())
	r.LoadHTMLGlob("templates/*")

	r.GET("/health", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "网页元数据分析",
			"pattern": string(cfg.GetPattern()),
		})
	})

	r.POST("/analyze", func(c *gin.Context) {
		document := c.PostForm("document")
		expr := c.PostForm("expr")
		timeStart := time.Now().UnixNano()

		dt, err := services.GetMetadataByData(document, expr)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error":    err.Error(),
				"TimeCost": time.Now().UnixNano() - timeStart,
				"Data":     "",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"error":    "",
			"TimeCost": time.Now().UnixNano() - timeStart,
			"Data":     dt,
		})
	})

	return r
}
