package app

import (
	"github.com/gin-gonic/gin"
	"your_spotify/internal/app/bean"
	"your_spotify/pkg/out"
)

func Web() {
	g := gin.Default()

	out.Table(bean.GetConfig(), "config")

	var basePath = bean.GetConfig().BaseURL()
	if basePath != "" && basePath[0] == '/' {
		basePath = basePath[1:]
	}

	route := g.Group(basePath)
	ApiRoute(route)

	err := g.Run(":9000")
	if err != nil {
		panic(err)
	}
}
