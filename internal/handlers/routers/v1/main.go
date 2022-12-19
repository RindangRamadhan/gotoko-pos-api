package v1

import (
	"gotoko-pos-api/internal"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type (
	RoutePath struct {
		Method   string
		Group    string
		Pattern  string
		Handlers []gin.HandlerFunc
	}
)

func Paths(container *internal.Container) []RoutePath {
	return []RoutePath{
		{
			Method: "GET", Group: "", Pattern: "/docs/*any", Handlers: []gin.HandlerFunc{
				ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1), ginSwagger.InstanceName("v1")),
			},
		},
	}

}
