package cashier

import (
	"gotoko-pos-api/common/request"
	"gotoko-pos-api/common/response"
	"gotoko-pos-api/common/validator"
	"gotoko-pos-api/internal/usecases/cashier/createcashier"

	"github.com/gin-gonic/gin"
)

// CreateCashierHandler godoc
// @Summary   Create a new cashier
// @Tags      Cashier
// @Security  JWTBearer
// @Accept    json
// @Produce   json
// @Param     payload  body      createcashier.InportRequest  true  "Payload"
// @Success   200      {object}  response.BodySuccess{data=createcashier.InportResponse}
// @Failure   500      {object}  response.BodyFailure{}
// @Router    /cashiers [post]
func CreateCashierHandler(inport createcashier.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createcashier.InportRequest

		// ? Binding Request
		if err := request.UnmarshalJSON(c, &req); err != nil {
			return
		}

		// ? Validate Request
		if err := validator.Validate(&req); err != nil {
			response.WriteError(c, "unable to validate payload", err)
			return
		}

		resp, err := inport.Execute(c.Copy().Request.Context(), req)
		if err != nil {
			response.WriteError(c, "failed when executing create cashier interactor", err)
			return
		}

		response.WriteSuccess(c, "Success", resp)
	}
}
