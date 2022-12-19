package server

import (
	"log"
	"net/http"
	"os"

	"gotoko-pos-api/common/middleware"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"

	cors "github.com/rs/cors/wrapper/gin"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
)

type GinHttpHandler struct {
	*GracefulShutdown
	Router *gin.Engine
}

func NewGinHttpRouter(address string) (*GinHttpHandler, error) {
	switch os.Getenv("ENV") {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	rate, err := limiter.NewRateFromFormatted("100-M")
	if err != nil {
		log.Fatal(err)
	}

	store := memory.NewStore()

	// TODO : Using Tracer

	router.Use(cors.New(
		cors.Options{
			AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodDelete},
			AllowedOrigins:   []string{"*"},
			AllowCredentials: true,
			AllowedHeaders:   []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept", "X-App-Token", "X-Requested-With", "Authorization"},
		},
	))

	// router.Use(apmgin.Middleware(router, apmgin.WithTracer(tracer)))
	router.Use(requestid.New())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(middleware.TDRLog())
	router.Use(mgin.NewMiddleware(limiter.New(store, rate, limiter.WithTrustForwardHeader(true))))

	return &GinHttpHandler{
		GracefulShutdown: NewGracefulShutdown(router, ":"+address),
		Router:           router,
	}, nil
}

func (g *GinHttpHandler) Start(address string) {
	g.GracefulShutdown.GracefullyShutdown()
}
