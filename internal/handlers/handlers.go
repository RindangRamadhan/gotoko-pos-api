package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gotoko-pos-api/docs"
	"gotoko-pos-api/internal"
	"gotoko-pos-api/internal/pkg/env"

	"gotoko-pos-api/internal/handlers/health"
	v1 "gotoko-pos-api/internal/handlers/routers/v1"
)

type (
	Router struct {
		router    *gin.Engine
		container *internal.Container
	}

	RoutePath struct {
		Method   string
		Group    string
		Pattern  string
		Handlers []gin.HandlerFunc
	}
)

func NewRouter(router *gin.Engine, container *internal.Container) *Router {
	setSwaggerGeneralInformation()

	return &Router{
		router:    router,
		container: container,
	}
}

func setSwaggerGeneralInformation() {
	docs.SwaggerInfo.Title = "GoToko POS API"
	docs.SwaggerInfo.Description = "This is GoToko POS API documentation."
	docs.SwaggerInfo.Version = env.Get().ServiceVersion
	docs.SwaggerInfo.Host = env.Get().ServiceURL
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func (_h *Router) RegisterRouter() {
	// Health Check
	_h.router.GET("/", health.HealthCheckHandler(_h.container.HealthCheckUsecase))

	// API V1
	apiv1 := _h.router.Group("/api/v1")
	for _, v := range v1.Paths(_h.container) {
		apiGroup := apiv1.Group(v.Group)

		switch v.Method {
		case http.MethodGet:
			apiGroup.GET(v.Pattern, v.Handlers...)
		case http.MethodPost:
			apiGroup.POST(v.Pattern, v.Handlers...)
		case http.MethodPut:
			apiGroup.PUT(v.Pattern, v.Handlers...)
		case http.MethodDelete:
			apiGroup.DELETE(v.Pattern, v.Handlers...)
		}
	}
}
