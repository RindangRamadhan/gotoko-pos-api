package health

import (
	"gotoko-pos-api/common/response"
	"gotoko-pos-api/internal/usecases/health"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler godoc
// @Summary  Health Check
// @Tags     Health-Check
// @Accept   json
// @Produce  json
// @Success  200  {object}  response.BodySuccess{data=health.InportResponse}
// @Failure  500  {object}  response.BodySuccess
// @Router   / [get]
func HealthCheckHandler(inport health.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := inport.Execute(c.Copy().Request.Context())
		response.WriteSuccess(c, "Success health check", resp, nil)
	}
}
