package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kooksee/html_meta/internal/config"
	"github.com/kooksee/html_meta/internal/kts"
	"github.com/kooksee/html_meta/internal/utils"
	"net/http"
)

func App() *gin.Engine {
	cfg := config.DefaultConfig()
	cfg.Init()
	kts.Init()

	r := gin.Default()
	r.GET("/health", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})

	r.POST("/parser/:name", func(context *gin.Context) {
		hp := &kts.HtmlPattern{Name: context.Param("name")}
		if err := hp.GetPattern(); err != nil {
			context.String(http.StatusBadRequest, err.Error())
			return
		}

		dt, err := context.GetRawData()
		if err != nil {
			context.String(http.StatusBadRequest, err.Error())
			return
		}

		ret, err := utils.UnMashallHtml(dt, hp.Pattern)
		if err != nil {
			context.String(http.StatusBadRequest, err.Error())
		}

		context.JSON(http.StatusOK, ret)
	})

	r.GET("/patterns", func(context *gin.Context) {
		hp := &kts.HtmlPattern{}
		dt, err := hp.GetPatternNames()
		if err != nil {
			context.String(http.StatusBadRequest, err.Error())
			return
		}
		context.JSON(http.StatusOK, dt)
	})

	r.GET("/patterns/:name", func(context *gin.Context) {
		hp := &kts.HtmlPattern{Name: context.Param("name")}
		if err := hp.GetPattern(); err != nil {
			context.String(http.StatusBadRequest, err.Error())
			return
		}
		context.String(http.StatusOK, hp.Pattern)
	})

	return r
}
